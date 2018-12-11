package eu.adamwilliams.reftypes.prototype.ast;

import java.util.List;

public class FunctionCallStatement extends Statement {
    private FunctionCallExpression expr;

    public FunctionCallStatement(FunctionCallExpression expr) {
        this.expr = expr;
    }

    public FunctionCallExpression getExpr() {
        return expr;
    }
}
