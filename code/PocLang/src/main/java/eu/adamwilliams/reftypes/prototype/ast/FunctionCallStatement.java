package eu.adamwilliams.reftypes.prototype.ast;

import java.util.Optional;

public class FunctionCallStatement extends Statement {
    private FunctionCallExpression expr;

    public FunctionCallStatement(FunctionCallExpression expr) {
        this.expr = expr;
    }

    public FunctionCallExpression getExpr() {
        return expr;
    }

    @Override
    public Optional<Expression> execute() {
        this.expr.evaluate();
        return Optional.empty();
    }
}
