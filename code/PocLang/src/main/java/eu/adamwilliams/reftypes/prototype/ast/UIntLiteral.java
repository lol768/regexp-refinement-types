package eu.adamwilliams.reftypes.prototype.ast;

public class UIntLiteral {
    private long value;

    public UIntLiteral(long value) {
        this.value = value;
    }

    public long getValue() {
        return value;
    }
}
