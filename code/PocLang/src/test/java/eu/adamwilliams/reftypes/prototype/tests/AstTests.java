package eu.adamwilliams.reftypes.prototype.tests;

import eu.adamwilliams.reftypes.prototype.Application;
import eu.adamwilliams.reftypes.prototype.ErrorReport;
import eu.adamwilliams.reftypes.prototype.ast.*;
import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeCheckResults;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;
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

public class AstTests {

    @Test
    public void testMath() {
        BinaryOperationExpression binOpAdd = new BinaryOperationExpression(
                BinaryOperationType.INT_ADD,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(7L, (long) ((Long) binOpAdd.evaluate()));

        BinaryOperationExpression binOpSub = new BinaryOperationExpression(
                BinaryOperationType.INT_SUBTRACT,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(1L, (long) ((Long) binOpSub.evaluate()));

        BinaryOperationExpression binOpMultiply = new BinaryOperationExpression(
                BinaryOperationType.INT_MULTIPLY,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(12L, (long) ((Long) binOpMultiply.evaluate()));

        BinaryOperationExpression binOpDivide = new BinaryOperationExpression(
                BinaryOperationType.INT_DIVIDE,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(2, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(2L, (long) ((Long) binOpDivide.evaluate()));
    }

    @Test
    public void testStringConcat() {
        BinaryOperationExpression binOpConcat = new BinaryOperationExpression(
                BinaryOperationType.STRING_CONCAT,
                new StringLiteral("a", new TypeContainer(Type.STRING, null)),
                new StringLiteral("b", new TypeContainer(Type.STRING, null))
        );
        Assert.assertEquals("ab", binOpConcat.evaluate());
    }

    @Test
    public void testMathMoreComplex() {
        BinaryOperationExpression binOpSubEx = new BinaryOperationExpression(
                BinaryOperationType.INT_ADD,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );

        BinaryOperationExpression binOpEx = new BinaryOperationExpression(
                BinaryOperationType.INT_MULTIPLY,
                binOpSubEx,
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(21L, (long) ((Long) binOpEx.evaluate()));
    }

    @Test
    public void testEvaluateReturn() {
        String moreAdvancedProgram = "function Test(id: uint[> 1]): uint {\n" +
                "    return 1+1\n" +
                "}";

        ParseTree tree = getParseTree(moreAdvancedProgram);
        Application app = new Application();
        TypeCheckResults typeCheckResults = app.doTypeChecks(tree);
        FunctionDeclaration function = typeCheckResults.getFunctionTable().getFunctionByIdentifier("Test");

        ReturnStatement returnStmt = (ReturnStatement) function.getBody().getStatements().get(0);
        Assert.assertTrue(returnStmt.getValue().isPresent());
        long result = (long) returnStmt.getValue().get().evaluate();
        Assert.assertEquals(2, result);
    }

    private ParseTree getParseTree(String basicProgram) {
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
