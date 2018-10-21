package eu.adamwilliams.reftypes.prototype.tests;

import eu.adamwilliams.reftypes.prototype.Application;
import eu.adamwilliams.reftypes.prototype.parser.PocLangLexer;
import eu.adamwilliams.reftypes.prototype.parser.PocLangParser;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.ATNConfigSet;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.tree.ParseTree;
import org.junit.Assert;
import org.junit.Test;

import java.util.BitSet;

public class ParsingTests {

    @Test
    public void testBasicProgram() {
        String basicProgram = "function LookupUserById(id: uint[> 1]): void {\n" +
                "    return 1+1\n" +
                "}";

        ParseTree tree = getParseTree(basicProgram);
        Assert.assertTrue(tree.getText().contains("LookupUserById"));
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
    public void testMissingReturnCall() {
        String moreAdvancedProgram = "function LookupUserById(id: uint[> 1]): uint {\n" +
                "    CallToUnrelatedFunction(1+1)\n" +
                "}";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        Assert.assertTrue(app.doTypeChecks(tree).getReports().stream().anyMatch((errorReport -> errorReport.getMsg().contains("No return statement"))));
    }


    private ParseTree getParseTree(String basicProgram) {
        PocLangLexer lexer = new PocLangLexer(CharStreams.fromString(basicProgram));
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        PocLangParser parser = new PocLangParser(tokens);
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
