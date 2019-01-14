package eu.adamwilliams.reftypes.prototype.tests;

import eu.adamwilliams.reftypes.prototype.Application;
import eu.adamwilliams.reftypes.prototype.ErrorReport;
import eu.adamwilliams.reftypes.prototype.ast.FunctionDeclaration;
import eu.adamwilliams.reftypes.prototype.domain.TypeCheckResults;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import eu.adamwilliams.reftypes.prototype.parser.PocLex;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.ATNConfigSet;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.tree.ParseTree;
import org.junit.Assert;
import org.junit.Test;

import java.util.BitSet;
import java.util.List;

public class ParsingTests {

    @Test
    public void testBasicProgram() {
        String basicProgram = "function LookupUserById(id: uint[> 1]): void {\n" +
                "    return\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        TypeCheckResults results = app.doTypeChecks(tree);
        Assert.assertEquals(0, results.getReports().size());
    }

    @Test
    public void testIllegalVariableReference() {
        String basicProgram = "function LookupUserById(id: uint[> 1]): string {\n" +
                "    return data\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        Assert.assertTrue( app.doTypeChecks(tree).getReports().size() > 1);
    }

    @Test
    public void testVariableShadowsExisting() {
        String basicProgram = "function LookupUserById(id: uint[> 1]): string {\n" +
                "    var a: uint\n" +
                "    if (true) {\n" +
                "        var a: string\n" +
                "        return a\n" +
                "    }\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        List<ErrorReport> reports = app.doTypeChecks(tree).getReports();
        Assert.assertEquals(1, reports.size());
        Assert.assertTrue(reports.stream().anyMatch(rp -> rp.getMsg().contains("shadow")));
    }

    @Test
    public void testBoolType() {
        String basicProgram = "function LookupUserById(id: uint[> 1]): bool {\n" +
                "    return true\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        Assert.assertEquals(0, app.doTypeChecks(tree).getReports().size());
    }

    @Test
    public void testParseRegex() {
        String basicProgram = "function LookupUserById(str: string[/(goose)+/]): void {\n" +
                "    return\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        Assert.assertEquals(0, app.doTypeChecks(tree).getReports().size());
    }

    @Test
    public void testRegexViolation() {
        String basicProgram = "function LookupUserById(): string[/g+/] {\n" +
                "    return \"f\"\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        Assert.assertNotEquals(0, app.doTypeChecks(tree).getReports().size());
    }

    @Test
    public void testDotMatchesAny() {
        String basicProgram = "function LookupUserById(): string[/.+/] {\n" +
                "    return \"f\"\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        Assert.assertEquals(0, app.doTypeChecks(tree).getReports().size());
    }

    @Test
    public void testNestedIf() {
        String basicProgram = "function LookupUserById(): string[/.+/] {\n" +
                "    if (true) {\n" +
                "        if (false) {\n" +
                "            return \"foo\"\n" +
                "        }\n" +
                "    }\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        Assert.assertEquals(0, app.doTypeChecks(tree).getReports().size());
    }

    @Test
    public void testIllegalIf() {
        String basicProgram = "function LookupUserById(): string[/.+/] {\n" +
                "    if (\"true\") {\n" +
                "        return \"it is true\"\n" +
                "    }\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        Assert.assertEquals(1, app.doTypeChecks(tree).getReports().size());
    }

    @Test
    public void testIllegalElseIf() {
        String basicProgram = "function LookupUserById(): string[/.+/] {\n" +
                "    if (true) {\n" +
                "        return \"it is true\"\n" +
                "    } else if (\"hmm\") {\n"+
                "        return \"abc\"\n" +
                "    }\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        TypeCheckResults results = app.doTypeChecks(tree);
        Assert.assertEquals(1, results.getReports().size());
        FunctionDeclaration lookupUserById = results.getFunctionTable().getFunctionByIdentifier("LookupUserById");
        Assert.assertNotNull(lookupUserById);
    }

    @Test
    public void testRegexAlternation() {
        String basicProgram = "function ShellFunction(): string[/g+|f+/] {\n" +
                "    return \"f\"\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Application app = new Application();
        TypeCheckResults result = app.doTypeChecks(tree);
        Assert.assertEquals(0, result.getReports().size());
    }

    @Test
    public void testRegexFunctionCallViolation() {
        String basicProgram = "function ShellFunction(): string[/g+|f+/] {\n" +
                "    return \"f\"\n" +
                "}\n" +
                "function Main(): string[/f+/] {\n" +
                "    return ShellFunction()\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Application app = new Application();
        // there should be a violation found
        Assert.assertNotEquals(0, app.doTypeChecks(tree).getReports().size());
    }

    @Test
    public void testBasicProgramWithViolatedUintConstraint() {
        String basicProgram = "function LookupUserById(id: uint[< 5]): void {\n" +
                "    return 5\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
        Application app = new Application();
        List<ErrorReport> reports = app.doTypeChecks(tree).getReports();
        Assert.assertTrue(reports.stream().anyMatch(p -> p.getMsg().contains("didn't satisfy")));
    }

    @Test
    public void testFunctionRedeclaration() {
        String moreAdvancedProgram = "function LookupUserById(id: uint[> 1]): void {\n" +
                "    return 1+1\n" +
                "}\n" + "function LookupUserById(id: uint[> 1]): uint {\n" +
                "    return 1+1\n" +
                "}";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("Redeclaration"))));
    }

    @Test
    public void testReturnFromVoid() {
        String moreAdvancedProgram = "function LookupUserById(id: uint): void {\n" +
                "    return 1\n" +
                "}\n";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("Attempt to return value from void function"))));
    }

    @Test
    public void testMissingReturnCall() {
        String moreAdvancedProgram = "function LookupUserById(id: uint[> 1]): uint {\n" +
                "    CallToUnrelatedFunction(1+1)\n" +
                "}";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("No return statement"))));
    }

    @Test
    public void testFunctionCallTypes() {
        String moreAdvancedProgram = "function LookupUserById(id: uint[> 1]): void {\n" +
                "    return 1+1\n" +
                "}\n" + "function Main(id: uint[> 1]): uint {\n" +
                "    return LookupUserById(\"5\")\n" +
                "}";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("didn't satisfy"))));
    }

    @Test
    public void testNestedFunctionCallTypes() {
        String moreAdvancedProgram = "function LookupUserById(id: uint[> 1]): uint[> 1] {\n" +
                "    return 1+1\n" +
                "}\n" + "function Main(): uint[< 5] {\n" +
                "    return LookupUserById(LookupUserById(2))\n" +
                "}";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("didn't satisfy"))));
    }

    @Test
    public void testMightViolateConstraint() {
        String moreAdvancedProgram = "function LookupUserById(id: uint[> 5]): uint {\n" +
                "    return 1+1\n" +
                "}\n" + "function Main(a: uint[> 1]): void {\n" +
                "    LookupUserById(a)\n" +
                "}";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("didn't satisfy"))));
    }

    @Test
    public void testFunctionCallTypesUsingVariable() {
        String moreAdvancedProgram = "function LookupUserById(id: uint[> 1]): void {\n" +
                "    return 1+1\n" +
                "}\n" + "function Main(id: uint[> 1]): uint {\n" +
                "    var a: string\n" +
                "    return LookupUserById(a)\n" +
                "}";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("didn't satisfy"))));
    }

    @Test
    public void testNestedFunctionCalls() {
        String nestedFunctionCalls = "function LookupUserById(id: uint[> 1]): void {\n" +
                "    return 1+1\n" +
                "}\n" + "function Main(id: uint[> 1]): uint {\n" +
                "    return LookupUserById(GetString())\n" +
                "}\n" + "function GetString(): string {\n" +
                "    return \"a\"\n" +
                "}";

        ParseTree tree = getParseTree(nestedFunctionCalls);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("didn't satisfy"))));
    }

    @Test
    public void testIllegalConstraintApplication() {
        String illegalConstraint = "function LookupUserById(id: string[> 1]): void {\n" +
                "    return 1+1\n" +
                "}\n";

        ParseTree tree = getParseTree(illegalConstraint);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("Attempt to apply int constraint to string type"))));
    }

    public static ParseTree getParseTree(String basicProgram) {
        PocLex lexer = new PocLex(CharStreams.fromString(basicProgram));
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        PocLang parser = new PocLang(tokens);

        parser.addErrorListener(new ANTLRErrorListener() {
            @Override
            public void syntaxError(Recognizer<?, ?> recognizer, Object offendingSymbol, int line, int charPositionInLine, String msg, RecognitionException e) {
                Assert.fail(msg);
            }

            @Override
            public void reportAmbiguity(Parser recognizer, DFA dfa, int startIndex, int stopIndex, boolean exact, BitSet ambigAlts, ATNConfigSet configs) {

            }

            @Override
            public void reportAttemptingFullContext(Parser recognizer, DFA dfa, int startIndex, int stopIndex, BitSet conflictingAlts, ATNConfigSet configs) {

            }

            @Override
            public void reportContextSensitivity(Parser recognizer, DFA dfa, int startIndex, int stopIndex, int prediction, ATNConfigSet configs) {

            }
        });
        return parser.program();
    }

}
