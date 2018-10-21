package eu.adamwilliams.reftypes.prototype.domain;

public class TypeContainer {

    private Type type;
    private String refinement; // TODO: better type representation

    public TypeContainer(Type type, String refinement) {
        this.type = type;
        this.refinement = refinement;
    }

    public Type getType() {
        return type;
    }

    public String getRefinement() {
        return refinement;
    }
}
