package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public class UIntLiteral extends ValueExpression {
    private long value;

    public UIntLiteral(long value, TypeContainer tc) {
        this.value = value;
        this.typeContainer = tc;
        if (this.typeContainer.getType() != Type.UNSIGNED_INTEGER) {
            throw new IllegalArgumentException("Illegal AST node, UInt literal must be UNSIGNED_INTEGER");
        }
    }

    @Override
    public Object getValue() {
        return value;
    }
}
