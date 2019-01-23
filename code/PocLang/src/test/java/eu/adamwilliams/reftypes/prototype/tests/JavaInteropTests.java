package eu.adamwilliams.reftypes.prototype.tests;

import eu.adamwilliams.reftypes.prototype.ast.Expression;
import eu.adamwilliams.reftypes.prototype.ast.JavaCallExpression;
import eu.adamwilliams.reftypes.prototype.ast.StringLiteral;
import eu.adamwilliams.reftypes.prototype.ast.UIntLiteral;
import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;
import org.junit.Test;

import java.util.ArrayList;

public class JavaInteropTests {

    @Test
    public void javaFunctionCall() {
        ArrayList<Expression> exprs = new ArrayList<>();
        exprs.add(new UIntLiteral(6, new TypeContainer(Type.UNSIGNED_INTEGER, null)));
        JavaCallExpression jce = new JavaCallExpression(
                "java.lang.System.out.println".split("\\."), exprs);
        jce.evaluate();

        exprs = new ArrayList<>();
        exprs.add(new StringLiteral("Hello world", new TypeContainer(Type.STRING, null)));
        jce = new JavaCallExpression(
                "java.lang.System.out.println".split("\\."), exprs);
        jce.evaluate();

    }
}
