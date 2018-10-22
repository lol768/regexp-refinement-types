package eu.adamwilliams.reftypes.prototype;

import eu.adamwilliams.reftypes.prototype.domain.VisitorPhase;
import eu.adamwilliams.reftypes.prototype.parser.PocLangLexer;
import eu.adamwilliams.reftypes.prototype.parser.PocLangParser;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTree;
import org.antlr.v4.runtime.tree.ParseTreeWalker;
import org.fusesource.jansi.Ansi;
import org.fusesource.jansi.AnsiConsole;
import org.sosy_lab.common.configuration.InvalidConfigurationException;
import org.sosy_lab.java_smt.SolverContextFactory;
import org.sosy_lab.java_smt.api.*;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.stream.Collectors;

public class Application {

    public static void main(String[] args) throws InterruptedException, SolverException, InvalidConfigurationException {
        AnsiConsole.systemInstall();
        System.out.println(Ansi.ansi().fg(Ansi.Color.GREEN)+"Reading program from stdin (use Ctrl+D when finished)..." + Ansi.ansi().fg(Ansi.Color.RED));
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
            System.err.printf(Ansi.ansi().fg(Ansi.Color.RED)+"%d syntax error(s) need to be resolved.%n", parser.getNumberOfSyntaxErrors());
            System.exit(-1);
        }

        ErrorReporter errorReporter = null;
        try {
            errorReporter = this.doTypeChecks(tree);
        } catch (InvalidConfigurationException e) {
            e.printStackTrace();
            System.exit(-1);
        }

        // Print out any and all errors
        for (ErrorReport report : errorReporter.getReports()) {
            System.err.println(Ansi.ansi().fg(Ansi.Color.RED)+"L" + report.getToken().getLine() + ":" + report.getToken().getCharPositionInLine() + " " + report.getMsg());
        }

        if (errorReporter.getReports().size() == 0) {
            System.out.println(Ansi.ansi().fg(Ansi.Color.GREEN)+"No errors to report");
        }
    }

    public ErrorReporter doTypeChecks(ParseTree tree) throws InvalidConfigurationException {
        ParseTreeWalker walker = new ParseTreeWalker();

        FunctionTable tableForProgram = new FunctionTable();
        ErrorReporter reporter = new ErrorReporterImpl();
        try (SolverContext context = SolverContextFactory.createSolverContext(SolverContextFactory.Solvers.SMTINTERPOL)) {
            VisitorListener listener = new VisitorListener(tableForProgram, reporter, context);
            walker.walk(listener, tree);

            // Now that our function table is populated we can check the function calls
            listener.setPhase(VisitorPhase.CHECKING_TYPES);
            walker.walk(listener, tree);
        }
        // We walk the parse tree in two passes
        // In the first phase, we're simply making a note of all of the declared functions in the function table

        return reporter;
    }


    private String readStdIn() {
        return new BufferedReader(new InputStreamReader(System.in))
                .lines().collect(Collectors.joining("\n"));
    }

}
