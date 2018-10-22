package eu.adamwilliams.reftypes.prototype.domain;

import eu.adamwilliams.reftypes.prototype.parser.PocLangParser;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

public class FunctionDeclaration {
    private String identifier;
    private TypeContainer returnType;
    private LinkedHashMap<String, TypeContainer> arguments;
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

    public LinkedHashMap<String, TypeContainer> getArguments() {
        return arguments;
    }

    public Map.Entry<String, TypeContainer> getNthArgument(int i) {
        List<Map.Entry<String, TypeContainer>> entries = new ArrayList<>(this.arguments.entrySet());
        return entries.get(i);
    }

    public FunctionDeclaration(String identifier, TypeContainer returnType, LinkedHashMap<String, TypeContainer> arguments, int lineNo, int colNo) {
        this.identifier = identifier;
        this.returnType = returnType;
        this.lineNo = lineNo;
        this.colNo = colNo;
        this.arguments = arguments;
    }
}
