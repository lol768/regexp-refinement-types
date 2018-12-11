package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public class ErrorNode extends ValueExpression {
    public ErrorNode() {
        this.typeContainer = new TypeContainer(Type.STRING, null);
    }

    @Override
    public String getValue() {
        return "Error";
    }
}
