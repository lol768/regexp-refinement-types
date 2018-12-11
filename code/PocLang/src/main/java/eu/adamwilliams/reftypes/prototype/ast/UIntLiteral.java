package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public class UIntLiteral extends ValueExpression {
    private long value;

    public UIntLiteral(long value, TypeContainer tc) {
        this.value = value;
        this.typeContainer = tc;
    }

    @Override
    public Object getValue() {
        return value;
    }
}
