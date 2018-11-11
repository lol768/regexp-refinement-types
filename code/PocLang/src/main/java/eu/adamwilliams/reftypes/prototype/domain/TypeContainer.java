package eu.adamwilliams.reftypes.prototype.domain;

import com.microsoft.z3.BoolExpr;

public class TypeContainer {

    private Type type;
    private BoolExpr refinement;

    public TypeContainer(Type type, BoolExpr refinement) {
        this.type = type;
        this.refinement = refinement;
    }

    public Type getType() {
        return type;
    }

    public BoolExpr getRefinement() {
        return refinement;
    }

    @Override
    public String toString() {
        String friendlyType = "";
        switch (this.type) {

            case UNSIGNED_INTEGER:
                friendlyType = "uint";
                break;
            case STRING:
                friendlyType = "string";
                break;
            case VOID:
                friendlyType = "void";
                break;
        }
        return friendlyType + (this.refinement == null ? "" : String.format(" [%s]", this.refinement.toString()));
    }
}
