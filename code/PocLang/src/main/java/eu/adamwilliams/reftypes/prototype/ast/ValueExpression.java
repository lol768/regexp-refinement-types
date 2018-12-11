package eu.adamwilliams.reftypes.prototype.ast;

public abstract class ValueExpression extends Expression {

    public abstract Object getValue();

    @Override
    public Object evaluate() {
        return this.getValue();
    }
}
