package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public class StringLiteral extends ValueExpression {
    private String value;

    public StringLiteral(String value, TypeContainer tc) {
        this.value = value;
        this.typeContainer = tc;
    }

    @Override
    public String getValue() {
        return value;
    }
}
