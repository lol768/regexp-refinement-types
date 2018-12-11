package eu.adamwilliams.reftypes.prototype.ast;

public class VariableDeclarationStatement extends Statement {
    private String identifier;

    public VariableDeclarationStatement(String identifier) {
        this.identifier = identifier;
    }

    public String getIdentifier() {
        return identifier;
    }
}
