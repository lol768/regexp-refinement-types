package eu.adamwilliams.reftypes.prototype;

import com.microsoft.z3.*;
import eu.adamwilliams.reftypes.prototype.domain.VisitorPhase;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import eu.adamwilliams.reftypes.prototype.parser.PocLex;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.RuleContext;
import org.antlr.v4.runtime.tree.ParseTree;
import org.antlr.v4.runtime.tree.ParseTreeWalker;
import org.fusesource.jansi.Ansi;
import org.fusesource.jansi.AnsiConsole;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.HashMap;
import java.util.Map;
import java.util.Stack;
import java.util.stream.Collectors;

public class Application {

    public static void main(String[] args) {
        AnsiConsole.systemInstall();
        System.out.println(Ansi.ansi().fg(Ansi.Color.GREEN) + "Reading program from stdin (use Ctrl+D when finished)..." + Ansi.ansi().fg(Ansi.Color.RED));
        Application instance = new Application();
        instance.handleProgram();
    }

    public void handleProgram() {
        String inputProgram = readStdIn();
        PocLex lexer = new PocLex(CharStreams.fromString(inputProgram));
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        PocLang parser = new PocLang(tokens);
        String[] ruleNames = parser.getRuleNames();
        ParseTree tree = parser.program();
        if (parser.getNumberOfSyntaxErrors() > 0) {
            // ANTLR is pretty forgiving, but generally we want to give up
            // if we get a syntactically invalid program for this prototype
            System.err.printf(Ansi.ansi().fg(Ansi.Color.RED) + "%d syntax error(s) need to be resolved.%n", parser.getNumberOfSyntaxErrors());
            System.exit(-1);
        }
        
        generateLatexDiagram(tree, ruleNames);

        ErrorReporter errorReporter = this.doTypeChecks(tree);

        // Print out any and all errors
        for (ErrorReport report : errorReporter.getReports()) {
            System.err.println(Ansi.ansi().fg(Ansi.Color.RED) + "L" + report.getToken().getLine() + ":" + report.getToken().getCharPositionInLine() + " " + report.getMsg());
        }

        if (errorReporter.getReports().size() == 0) {
            System.out.println(Ansi.ansi().fg(Ansi.Color.GREEN) + "No errors to report");
        }
    }

    private void generateLatexDiagram(ParseTree treeItem, String[] ruleNames) {
        if (treeItem.getPayload() instanceof RuleContext) {
            RuleContext rc = (RuleContext) treeItem.getPayload();
            System.out.print("\\node{"+ruleNames[rc.getRuleIndex()]+"}\n");
        }
        for (int i = 0; i < treeItem.getChildCount(); i++) {
            generateLatexDiagramAux(treeItem.getChild(i), 0, ruleNames);
        }


    }

    private void generateLatexDiagramAux(ParseTree treeItem, int indentation, String[] ruleNames) {
        StringBuilder indent = new StringBuilder();
        for (int i = 0; i < indentation; i++) {
            indent.append(" ");
        }
        if (treeItem.getPayload() instanceof RuleContext) {
            RuleContext rc = (RuleContext) treeItem.getPayload();
            String text = ruleNames[rc.getRuleIndex()];
            sanitiseText(indent, text);
        } else {
            String text = treeItem.getText();
            sanitiseText(indent, text);
        }
        if (treeItem.getChildCount() > 0) {
            System.out.print("\n"+indent.toString());
        }
        else {
            System.out.println("}");
            return;
        }
        for (int i = 0; i < treeItem.getChildCount(); i++) {
            generateLatexDiagramAux(treeItem.getChild(i), indentation+2, ruleNames);
        }
        System.out.println(indent.toString() + "}");
    }

    private void sanitiseText(StringBuilder indent, String text) {
        text = text.replace("_", "\\_");
        text = text.replace("{", "\\{");
        text = text.replace("}", "\\}");
        text = text.replace("\n", "\\n");
        System.out.print(indent.toString() + "child { node{"+text+"} ");
    }


    public ErrorReporter doTypeChecks(ParseTree tree) {
        ParseTreeWalker walker = new ParseTreeWalker();

        FunctionTable tableForProgram = new FunctionTable();
        ErrorReporter reporter = new ErrorReporterImpl();
        Map<String, String> cfg = new HashMap<>();
        cfg.put("model", "true");

        try (Context ctx = new Context(cfg)) {
            VisitorListener listener = new VisitorListener(tableForProgram, reporter, ctx);
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
