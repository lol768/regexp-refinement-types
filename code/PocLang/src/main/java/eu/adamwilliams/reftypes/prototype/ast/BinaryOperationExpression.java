package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

public class BinaryOperationExpression extends Expression {
    private BinaryOperationType binOpType;
    private Expression lhs;
    private Expression rhs;

    public BinaryOperationExpression(BinaryOperationType binOpType, Expression lhs, Expression rhs) {
        this.binOpType = binOpType;
        this.lhs = lhs;
        this.rhs = rhs;
        this.setType();
    }

    public BinaryOperationType getBinOpType() {
        return this.binOpType;
    }

    public Expression getLeft() {
        return this.lhs;
    }

    public Expression getRight() {
        return this.rhs;
    }

    private void setType() {
        switch (this.binOpType) {
            case INT_ADD:
            case INT_SUBTRACT:
            case INT_DIVIDE:
            case INT_MULTIPLY:
                if (this.lhs.getTypeContainer().getType() != Type.UNSIGNED_INTEGER || this.rhs.getTypeContainer().getType() != Type.UNSIGNED_INTEGER) {
                    throw new UnsupportedOperationException("Cannot do arithmetic on non-integers");
                }
                this.typeContainer = new TypeContainer(Type.UNSIGNED_INTEGER, null); // TODO: properly infer refinement
                break;
            case STRING_CONCAT:
                if (this.lhs.getTypeContainer().getType() != Type.STRING || this.rhs.getTypeContainer().getType() != Type.STRING) {
                    throw new UnsupportedOperationException("Cannot concatenate non-strings");
                }
                this.typeContainer = new TypeContainer(Type.STRING, null); // TODO: properly infer refinement
                break;
        }
    }

    @Override
    public Object evaluate() {
        // TODO: eval()
        return null;
    }
}
