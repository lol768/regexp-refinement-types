package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.StackEntry;

import java.util.Optional;

public class VariableAssignmentStatement extends Statement {
    private Expression newValue;
    private String variableRef;
    private StackEntry entry;

    public VariableAssignmentStatement(Expression value, String variableRef, StackEntry entry) {
        this.newValue = value;
        this.variableRef = variableRef;
        this.entry = entry;
    }

    public Expression getNewValue() {
        return newValue;
    }

    @Override
    public Optional<Expression> execute() {
        this.entry.setCurrentValue(this.newValue.evaluate());
        return Optional.empty();
    }
}
