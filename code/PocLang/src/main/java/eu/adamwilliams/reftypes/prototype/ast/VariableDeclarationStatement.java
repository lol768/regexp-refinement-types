package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.StackEntry;

public class VariableDeclarationStatement extends Statement {
    private String identifier;
    private StackEntry stackEntry;

    public VariableDeclarationStatement(String identifier, StackEntry stackEntry) {
        this.identifier = identifier;
        this.stackEntry = stackEntry;
    }

    public StackEntry getStackEntry() {
        return stackEntry;
    }

    public String getIdentifier() {
        return identifier;
    }

    @Override
    public void execute() {
        // no-op
    }
}
