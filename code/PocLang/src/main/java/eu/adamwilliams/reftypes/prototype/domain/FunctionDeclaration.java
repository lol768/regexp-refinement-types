package eu.adamwilliams.reftypes.prototype.domain;

public class FunctionDeclaration {
    private String identifier;
    private Type returnType;
    private int lineNo;
    private int colNo;

    public String getIdentifier() {
        return identifier;
    }

    public Type getReturnType() {
        return returnType;
    }

    public int getLineNo() {
        return lineNo;
    }

    public int getColNo() {
        return colNo;
    }

    public FunctionDeclaration(String identifier, Type returnType, int lineNo, int colNo) {
        this.identifier = identifier;
        this.returnType = returnType;
        this.lineNo = lineNo;
        this.colNo = colNo;
    }
}
