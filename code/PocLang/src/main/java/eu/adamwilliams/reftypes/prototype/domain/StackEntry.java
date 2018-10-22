package eu.adamwilliams.reftypes.prototype.domain;

public class StackEntry {
    private TypeContainer type;
    private StackEntryType stackType;

    public StackEntry(TypeContainer type, StackEntryType stackType) {
        this.type = type;
        this.stackType = stackType;
    }

    public TypeContainer getType() {
        return type;
    }

    public StackEntryType getStackType() {
        return stackType;
    }
}

