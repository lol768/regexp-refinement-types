package eu.adamwilliams.reftypes.prototype.domain;

import eu.adamwilliams.reftypes.prototype.ast.Body;

import java.util.HashMap;
import java.util.Map;

public class ScopeContainer {
    private Map<String, StackEntry> identifierMap = new HashMap<String, StackEntry>();
    private Body enclosingBody;
    public ScopeContainer(Body body) {
        this.enclosingBody = body;
    }

    public Body getEnclosingBody() {
        return enclosingBody;
    }

    public void insertIdentifier(String identifierName, StackEntry entry) {
        this.identifierMap.put(identifierName, entry);
    }

    public boolean hasIdentifier(String name) {
        return this.identifierMap.containsKey(name);
    }

    public StackEntry getByIdentifier(String name) {
        return this.identifierMap.get(name);
    }


}
