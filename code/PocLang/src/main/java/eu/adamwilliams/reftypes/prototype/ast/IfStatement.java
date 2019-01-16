package eu.adamwilliams.reftypes.prototype.ast;

import java.util.Optional;

public class IfStatement extends Statement {
    private Expression condition;
    private Body bodyIfTrue;
    private Body elseBody; // may be null

    private IfStatement elseIf; // may be null, these chain

    public IfStatement(Expression condition, Body bodyIfTrue, Body elseBody, IfStatement elseIf) {
        this.condition = condition;
        this.bodyIfTrue = bodyIfTrue;
        this.elseBody = elseBody;
        this.elseIf = elseIf;
    }

    public Expression getCondition() {
        return condition;
    }

    public Optional<Body> getBodyIfTrue() {
        return Optional.ofNullable(bodyIfTrue);
    }

    public void setBodyIfTrue(Body bodyIfTrue) {
        this.bodyIfTrue = bodyIfTrue;
    }

    public Optional<Body> getElseBody() {
        return Optional.ofNullable(elseBody);
    }

    public Optional<IfStatement> getElseIf() {
        return Optional.ofNullable(elseIf);
    }

    public void setElseBody(Body elseBody) {
        this.elseBody = elseBody;
    }

    public void setElseIf(IfStatement elseIf) {
        this.elseIf = elseIf;
    }

    @Override
    public Optional<Expression> execute() {
        boolean expr = (boolean) this.condition.evaluate();
        if (expr) {
            return BodyEvaluator.evaluateBody(this.bodyIfTrue);
        } else if (elseBody != null) {
            return BodyEvaluator.evaluateBody(this.elseBody);
        } else if (elseIf != null) {
            return elseIf.execute();
        }
        throw new IllegalArgumentException("Illegal if");
    }
}

