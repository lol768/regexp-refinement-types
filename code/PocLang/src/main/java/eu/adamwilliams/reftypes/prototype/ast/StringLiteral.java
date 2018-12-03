package eu.adamwilliams.reftypes.prototype.ast;

public class StringLiteral {
    private String value;

    public StringLiteral(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }
}
