package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public class StringLiteral extends ValueExpression {
    private String value;

    public StringLiteral(String value, TypeContainer tc) {
        this.value = value;
        this.typeContainer = tc;
        if (this.typeContainer.getType() != Type.STRING) {
            throw new IllegalArgumentException("Illegal AST node, String literal must be STRING");
        }
    }

    @Override
    public String getValue() {
        return value;
    }
}
