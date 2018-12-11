package eu.adamwilliams.reftypes.prototype;

import eu.adamwilliams.reftypes.prototype.ast.FunctionDeclaration;

import java.util.HashMap;
import java.util.Map;

public class FunctionTable {
    private Map<String, FunctionDeclaration> table = new HashMap<>();

    public FunctionDeclaration getFunctionByIdentifier(String identifier) {
        return this.table.get(identifier);
    }

    public boolean hasFunction(String identifier) {
        return this.table.containsKey(identifier);
    }

    public void addFunction(FunctionDeclaration declaration) {
        this.table.put(declaration.getIdentifier(), declaration);
    }

}
