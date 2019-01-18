package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public class ComparisonOperation extends Expression {
    private ComparisonOperationType comparisonType;
    private Expression lhs;
    private Expression rhs;

    public ComparisonOperation(ComparisonOperationType comparisonType, Expression lhs, Expression rhs) {
        this.comparisonType = comparisonType;
        this.lhs = lhs;
        this.rhs = rhs;
        this.typeContainer = new TypeContainer(Type.BOOLEAN, null);
    }

    public ComparisonOperationType getComparisonType() {
        return this.comparisonType;
    }

    public Expression getLeft() {
        return this.lhs;
    }

    public Expression getRight() {
        return this.rhs;
    }

    @Override
    public Object evaluate() {
        switch (this.comparisonType) {
            case EQUALS:
                return lhs.evaluate().equals(rhs.evaluate());
            case GT:
                return ((Long)lhs.evaluate()) > ((Long)rhs.evaluate());
            case GE:
                return ((Long)lhs.evaluate()) >= ((Long)rhs.evaluate());
            case LT:
                return ((Long)lhs.evaluate()) < ((Long)rhs.evaluate());
            case LE:
                return ((Long)lhs.evaluate()) <= ((Long)rhs.evaluate());
        }
        return null;
    }
}
