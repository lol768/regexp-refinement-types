package eu.adamwilliams.reftypes.prototype.domain;

public class StackEntry {
    private TypeContainer type;
    private final StackEntryType stackEntryType;
    private Object currentValue;

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

    public Object getCurrentValue() {
        return currentValue;
    }

    public void setCurrentValue(Object currentValue) {
        this.currentValue = currentValue;
    }
}

