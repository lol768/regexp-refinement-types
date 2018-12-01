package eu.adamwilliams.reftypes.prototype.domain;

public class StackEntry {
    private TypeContainer type;
    private final StackEntryType stackEntryType;

    public StackEntry(TypeContainer type, StackEntryType stackEntryType) {
        this.type = type;
        this.stackEntryType = stackEntryType;
    }

    public TypeContainer getType() {
        return type;
    }

    public StackEntryType getStackEntryType() {
        return stackEntryType;
    }
}

