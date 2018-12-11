package eu.adamwilliams.reftypes.prototype.tests;

import eu.adamwilliams.reftypes.prototype.Application;
import eu.adamwilliams.reftypes.prototype.ErrorReport;
import eu.adamwilliams.reftypes.prototype.ast.BinaryOperationExpression;
import eu.adamwilliams.reftypes.prototype.ast.BinaryOperationType;
import eu.adamwilliams.reftypes.prototype.ast.StringLiteral;
import eu.adamwilliams.reftypes.prototype.ast.UIntLiteral;
import eu.adamwilliams.reftypes.prototype.domain.Type;
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

}
