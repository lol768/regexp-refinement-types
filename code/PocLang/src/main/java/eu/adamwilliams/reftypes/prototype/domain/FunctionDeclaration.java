package eu.adamwilliams.reftypes.prototype.domain;

import java.util.Map;

public class FunctionDeclaration {
    private String identifier;
    private TypeContainer returnType;
    private Map<String, TypeContainer> arguments;
    private int lineNo;
    private int colNo;

    public String getIdentifier() {
        return identifier;
    }

    public TypeContainer getReturnType() {
        return returnType;
    }

    public int getLineNo() {
        return lineNo;
    }

    public int getColNo() {
        return colNo;
    }

    public FunctionDeclaration(String identifier, TypeContainer returnType, Map<String, TypeContainer> arguments, int lineNo, int colNo) {
        this.identifier = identifier;
        this.returnType = returnType;
        this.lineNo = lineNo;
        this.colNo = colNo;
        this.arguments = arguments;
    }
}