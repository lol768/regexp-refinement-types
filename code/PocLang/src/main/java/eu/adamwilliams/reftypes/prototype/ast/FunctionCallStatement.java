package eu.adamwilliams.reftypes.prototype.ast;

public class FunctionCallStatement extends Statement {
    private FunctionCallExpression expr;

    public FunctionCallStatement(FunctionCallExpression expr) {
        this.expr = expr;
    }

    public FunctionCallExpression getExpr() {
        return expr;
    }

    @Override
    public void execute() {
        expr.evaluate();
    }
}
