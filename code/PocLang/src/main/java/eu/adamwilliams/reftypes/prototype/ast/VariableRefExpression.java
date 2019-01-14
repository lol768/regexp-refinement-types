package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.StackEntry;

public class VariableRefExpression extends Expression {

    private String variableRef;
    private StackEntry entry;

    public VariableRefExpression(String variableRef, StackEntry entry) {
        this.variableRef = variableRef;
        this.entry = entry;
        this.typeContainer = this.entry.getType();
    }

    public String getVariableRef() {
        return variableRef;
    }

    public StackEntry getEntry() {
        return entry;
    }

    @Override
    public Object evaluate() {
        return this.entry.getCurrentValue();
    }
}
