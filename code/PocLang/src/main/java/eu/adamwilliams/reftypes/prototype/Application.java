package eu.adamwilliams.reftypes.prototype;

import com.microsoft.z3.*;
import eu.adamwilliams.reftypes.prototype.ast.BodyEvaluator;
import eu.adamwilliams.reftypes.prototype.ast.Expression;
import eu.adamwilliams.reftypes.prototype.domain.TypeCheckResults;
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
import spark.ResponseTransformer;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.*;
import java.util.stream.Collectors;

import com.google.gson.Gson;

import javax.json.Json;

import static spark.Spark.*;

public class Application {

    public static void main(String[] args) {
        AnsiConsole.systemInstall();
        System.out.println(Ansi.ansi().fg(Ansi.Color.GREEN) + "Reading program from stdin (use Ctrl+D when finished)..." + Ansi.ansi().fg(Ansi.Color.RED));
        Application instance = new Application();
        //instance.handleProgram();
        instance.startWebServer();
    }

    public void startWebServer() {
        Gson gson = new Gson();
        staticFiles.location("/public");
        before((request, response) -> response.type("application/json"));
        options("/check",
                (request, response) -> {

                    String accessControlRequestHeaders = request
                            .headers("Access-Control-Request-Headers");
                    if (accessControlRequestHeaders != null) {
                        response.header("Access-Control-Allow-Headers",
                                accessControlRequestHeaders);
                    }

                    String accessControlRequestMethod = request
                            .headers("Access-Control-Request-Method");
                    if (accessControlRequestMethod != null) {
                        response.header("Access-Control-Allow-Methods",
                                accessControlRequestMethod);
                    }

                    return "\"OK\"";
                });

        get("/language", (req, res) -> {
            return Files.readAllLines(
                    Paths.get(this.getClass().getResource("lang.json").toURI()), Charset.defaultCharset());
        });

        post("/run", (req, res) -> {
                    RiseForFunRequest riseForFunRequest = gson.fromJson(req.body(), RiseForFunRequest.class);
                    RiseForFunResponse response = new RiseForFunResponse();
                    response.setVersion(1.0);
                    ArrayList<Output> outputs = new ArrayList<>();
                    response.setOutputs(outputs);

                    String program = riseForFunRequest.getSource();
                    PocLex lexer = new PocLex(CharStreams.fromString(program));
                    CommonTokenStream tokens = new CommonTokenStream(lexer);
                    PocLang parser = new PocLang(tokens);
                    if (parser.getNumberOfSyntaxErrors() > 0) {
                        res.status(400);
                        outputs.add(new Output("text/plain", "Syntax error"));
                        return response;
                    }
                    List<ErrorReport> reports = doTypeChecks(parser.program()).getReports();
                    StringBuilder s = new StringBuilder();

                    int i = 0;
                    for (ErrorReport report : reports) {
                        i++;
                        s.append("snip(" +report.getLineNo()).append(",").append(report.getColNo()).append("): Error: ").append(report.getMsg()).append(".\n");
                    }

                    s.append("\nVerifier finished with " + i + " errors");

                    if (s.length() == 0) {
                        s.append("Type check succeeded");
                    }
                    outputs.add(new Output("text/plain", s.toString()));

                    return response;
                }, new JsonTransformer()
        );

        get("/metadata", (req, res) -> {
            RiseForFunMetadata metadata = new RiseForFunMetadata();
            metadata.setDescription("Simple language with refinement types supporting regex constraints");
            metadata.setName("RegexRefinementTypes");
            metadata.setDisplayName("Regex Refinement Types");
            metadata.setInstitutionUrl("https://warwick.ac.uk/services/its/");
            metadata.setInstitution("the University of Warwick");
            metadata.setInstitutionImageUrl("https://671da5df.ngrok.io/warwick.png?a");
            metadata.setTermsOfUseUrl("https://warwick.ac.uk/terms");
            metadata.setPrivacyUrl("https://warwick.ac.uk/terms/privacy/");
            metadata.setQuestion("Will the type-checker find a regex violation?");
            metadata.setVersion(1.0);
            metadata.setMimeType("text/rrt");
            metadata.setUrl("https://adamwilliams.eu");
            metadata.setDisableErrorTable(false);
            metadata.setSupportsLanguageSyntax(true);
            ArrayList<Sample> samples = new ArrayList<>();
            Sample e = new Sample();
            e.setName("Example 1: regex refinements");
            e.setSource("function Main(): string[/[a-f]+/] {\n" +
                    "    return Secondary()\n" +
                    "}\n" +
                    "\n" +
                    "function Secondary(): string[/[a-z][a-z]/] {\n" +
                    "    return \"ao\"\n" +
                    "}");
            samples.add(e);
            Sample f = new Sample();
            f.setName("Example 2: integer refinements");
            f.setSource("function Main(): string[/[a-f]+/] {\n" +
                    "    return Secondary()\n" +
                    "}\n" +
                    "\n" +
                    "function Secondary(): string[/[a-z][a-z]/] {\n" +
                    "    return \"ao\"\n" +
                    "}");
            samples.add(f);
            metadata.setSamples(samples);

            return metadata;
        }, new JsonTransformer());

        post("/check", (req, res) -> {
                    res.header("Access-Control-Allow-Origin", "*");
                    String program = req.body();
                    PocLex lexer = new PocLex(CharStreams.fromString(program));
                    CommonTokenStream tokens = new CommonTokenStream(lexer);
                    PocLang parser = new PocLang(tokens);
                    if (parser.getNumberOfSyntaxErrors() > 0) {
                        res.status(400);
                        return "Syntactically invalid";
                    }
                    return doTypeChecks(parser.program()).getReports();

                }, new JsonTransformer()
        );

        post("/evaluate", (req, res) -> {
                    res.header("Access-Control-Allow-Origin", "*");
                    CommonTokenStream tokens = new CommonTokenStream(new PocLex(CharStreams.fromString(req.body())));
                    PocLang parser = new PocLang(tokens);
                    if (parser.getNumberOfSyntaxErrors() > 0) {
                        res.status(400);
                        return "Syntactically invalid.";
                    }
                    TypeCheckResults results = doTypeChecks(parser.program());
                    if (!results.getReports().isEmpty()) {
                        res.status(400);
                        return "Type check failed, use /check endpoint.";
                    }

                    if (!results.getFunctionTable().hasFunction("Main")) {
                        res.status(400);
                        return "Need a Main function to execute.";
                    }

                    Optional<Expression> mainEvalResult = BodyEvaluator.evaluateBody(results.getFunctionTable().getFunctionByIdentifier("Main").getBody());
                    if (mainEvalResult.isPresent()) {
                        return mainEvalResult.get().evaluate();
                    } else {
                        return "Void method";
                    }

                }, new JsonTransformer()
        );
    }

    public void handleProgram() {
        String inputProgram = readStdIn();
        PocLex lexer = new PocLex(CharStreams.fromString(inputProgram));
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        PocLang parser = new PocLang(tokens);
        String[] ruleNames = parser.getRuleNames();
        //generateLatexDiagram(parser.type(), ruleNames);

        ParseTree tree = parser.program();
        if (parser.getNumberOfSyntaxErrors() > 0) {
            // ANTLR is pretty forgiving, but generally we want to give up
            // if we get a syntactically invalid program for this prototype
            System.err.printf(Ansi.ansi().fg(Ansi.Color.RED) + "%d syntax error(s) need to be resolved.%n", parser.getNumberOfSyntaxErrors());
            System.exit(-1);
        }


        TypeCheckResults tcr = this.doTypeChecks(tree);

        // Print out any and all errors
        for (ErrorReport report : tcr.getReports()) {
            System.err.println(Ansi.ansi().fg(Ansi.Color.RED) + "L" + report.getToken().getLine() + ":" + report.getToken().getCharPositionInLine() + " " + report.getMsg());
        }

        if (tcr.getReports().size() == 0) {
            System.out.println(Ansi.ansi().fg(Ansi.Color.GREEN) + "No errors to report");
        }
    }

    private void generateLatexDiagram(ParseTree treeItem, String[] ruleNames) {
        if (treeItem.getPayload() instanceof RuleContext) {
            RuleContext rc = (RuleContext) treeItem.getPayload();
            System.out.print("[" + ruleNames[rc.getRuleIndex()] + "\n");
        }
        for (int i = 0; i < treeItem.getChildCount(); i++) {
            generateLatexDiagramAux(treeItem.getChild(i), 0, ruleNames);
        }
        System.out.println("]");


    }

    private void generateLatexDiagramAux(ParseTree treeItem, int indentation, String[] ruleNames) {
        StringBuilder indent = new StringBuilder();
        for (int i = 0; i < indentation; i++) {
            indent.append(" ");
        }
        System.out.print("[");
        if (treeItem.getPayload() instanceof RuleContext) {
            RuleContext rc = (RuleContext) treeItem.getPayload();
            String text = ruleNames[rc.getRuleIndex()];
            System.out.print(sanitiseText(text));
        } else {
            String text = treeItem.getText();
            System.out.print(sanitiseText(text));
        }
        if (treeItem.getChildCount() > 0) {
            System.out.print("\n" + indent.toString());
        } else {
            System.out.println("]");
            return;
        }
        for (int i = 0; i < treeItem.getChildCount(); i++) {
            generateLatexDiagramAux(treeItem.getChild(i), indentation + 2, ruleNames);
        }
        System.out.println(indent.toString() + "]");
    }

    private String sanitiseText(String text) {
        text = text.replace("_", "\\_");
        text = text.replace("{", "\\{");
        text = text.replace("}", "\\}");
        text = text.replace("\n", "\\textbackslash{}n");
        if (text.contains("[") || text.contains("]")) {
            return "{" + text + "}";
        }
        return text;
    }


    public TypeCheckResults doTypeChecks(ParseTree tree) {
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
            return listener.getResults();
        }
        // We walk the parse tree in two passes
        // In the first phase, we're simply making a note of all of the declared functions in the function table
    }


    private String readStdIn() {
        return new BufferedReader(new InputStreamReader(System.in))
                .lines().collect(Collectors.joining("\n"));
    }

    public class JsonTransformer implements ResponseTransformer {

        private Gson gson = new Gson();

        @Override
        public String render(Object model) {
            return gson.toJson(model);
        }


    }
}
