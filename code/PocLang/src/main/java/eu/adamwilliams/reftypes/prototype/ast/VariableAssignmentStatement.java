package eu.adamwilliams.reftypes.prototype.ast;

public class VariableAssignmentStatement extends Statement {
    private Expression newValue;

    public VariableAssignmentStatement(Expression value) {
        this.newValue = value;
    }

    public Expression getNewValue() {
        return newValue;
    }
}
