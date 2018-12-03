package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public abstract class Expression {
    protected TypeContainer typeContainer;

    public TypeContainer getTypeContainer() {
        return typeContainer;
    }

    public abstract Object evaluate();
}
