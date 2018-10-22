package eu.adamwilliams.reftypes.prototype.domain;

import org.sosy_lab.java_smt.api.BooleanFormula;

public class TypeContainer {

    private Type type;
    private BooleanFormula refinement; // TODO: better type representation

    public TypeContainer(Type type, BooleanFormula refinement) {
        this.type = type;
        this.refinement = refinement;
    }

    public Type getType() {
        return type;
    }

    public BooleanFormula getRefinement() {
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
