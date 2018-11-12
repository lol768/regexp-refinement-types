package eu.adamwilliams.reftypes.prototype.tests;

import com.microsoft.z3.BoolExpr;
import com.microsoft.z3.Context;
import com.microsoft.z3.SeqExpr;
import eu.adamwilliams.reftypes.prototype.IRegexZ3Adapter;
import eu.adamwilliams.reftypes.prototype.RegexZ3Adapter;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import eu.adamwilliams.reftypes.prototype.parser.PocLex;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.ATNConfigSet;
import org.antlr.v4.runtime.dfa.DFA;
import org.junit.Assert;
import org.junit.Test;

import java.util.BitSet;
import java.util.HashMap;
import java.util.Map;

public class RegexZ3AdapterTests {

    @Test
    public void testSimpleExpression() {
        String constraint = "string[/a+/]";
        PocLang.TypeContext tree = getParseTree(constraint);
        Context z3Context = getZ3Context();
        IRegexZ3Adapter adapter = new RegexZ3Adapter();

        BoolExpr result = adapter.getRefinementType((SeqExpr) z3Context.mkConst("y", z3Context.getStringSort()), tree.string_constraint().re(), z3Context);
        Assert.assertEquals("(str.in.re y (re.+ (str.to.re \"a\")))", result.toString());
    }

    @Test
    public void testKleeneStarGrouped() {
        String constraint = "string[/(aaa)*/]";
        PocLang.TypeContext tree = getParseTree(constraint);
        Context z3Context = getZ3Context();
        IRegexZ3Adapter adapter = new RegexZ3Adapter();

        BoolExpr result = adapter.getRefinementType((SeqExpr) z3Context.mkConst("y", z3Context.getStringSort()), tree.string_constraint().re(), z3Context);
        Assert.assertEquals("(let ((a!1 (re.* (re.++ (str.to.re \"a\") (re.++ (str.to.re \"a\") (str.to.re \"a\"))))))\n" +
                "  (str.in.re y a!1))", result.toString());
    }

    private Context getZ3Context() {

        Map<String, String> cfg = new HashMap<>();
        cfg.put("model", "true");

        return new Context(cfg);
    }

    private PocLang.TypeContext getParseTree(String constraint) {
        PocLex lexer = new PocLex(CharStreams.fromString(constraint));
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
        return parser.type();
    }
}
