package eu.adamwilliams.reftypes.prototype;

import eu.adamwilliams.reftypes.prototype.domain.VisitorPhase;
import eu.adamwilliams.reftypes.prototype.parser.PocLangLexer;
import eu.adamwilliams.reftypes.prototype.parser.PocLangParser;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTree;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.stream.Collectors;

public class Application {

    public static void main(String[] args) {
        System.out.println("Reading program from stdin (use Ctrl+D when finished)...");
        Application instance = new Application();
        instance.handleProgram();
    }

    public void handleProgram() {
        String inputProgram = readStdIn();
        PocLangLexer lexer = new PocLangLexer(CharStreams.fromString(inputProgram));
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        PocLangParser parser = new PocLangParser(tokens);

        ParseTree tree = parser.program();
        if (parser.getNumberOfSyntaxErrors() > 0) {
            // ANTLR is pretty forgiving, but generally we want to give up
            // if we get a syntactically invalid program for this prototype
            System.err.printf("%d syntax error(s) need to be resolved.%n", parser.getNumberOfSyntaxErrors());
            System.exit(-1);
        }

        ErrorReporter errorReporter = this.doTypeChecks(tree);

        // Print out any and all errors
        for (ErrorReport report : errorReporter.getReports()) {
            System.err.println("L" + report.getToken().getLine() + ":" + report.getToken().getCharPositionInLine() + " " + report.getMsg());
        }

        if (errorReporter.getReports().size() == 0) {
            System.out.println("No errors to report");
        }
    }

    public ErrorReporter doTypeChecks(ParseTree tree) {
        ParseTreeWalker walker = new ParseTreeWalker();

        FunctionTable tableForProgram = new FunctionTable();
        ErrorReporter reporter = new ErrorReporterImpl();
        VisitorListener listener = new VisitorListener(tableForProgram, reporter);
        // We walk the parse tree in two passes
        // In the first phase, we're simply making a note of all of the declared functions in the function table
        walker.walk(listener, tree);

        // Now that our function table is populated we can check the function calls
        listener.setPhase(VisitorPhase.CHECKING_TYPES);
        walker.walk(listener, tree);
        return reporter;
    }


    private String readStdIn() {
        return new BufferedReader(new InputStreamReader(System.in))
                .lines().collect(Collectors.joining("\n"));
    }
}