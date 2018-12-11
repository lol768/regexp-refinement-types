package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public class BooleanLiteral extends ValueExpression {
    private boolean value;

    public BooleanLiteral(boolean value, TypeContainer tc) {
        this.value = value;
        this.typeContainer = tc;
        if (this.typeContainer.getType() != Type.STRING) {
            throw new IllegalArgumentException("Illegal AST node, Boolean literal must be BOOLEAN");
        }
    }

    @Override
    public Object getValue() {
        return this.value;
    }
}
