package eu.adamwilliams.reftypes.prototype.domain;

import java.util.HashMap;
import java.util.Map;

public class ScopeContainer {
    private Map<String, StackEntry> identifierMap = new HashMap<String, StackEntry>();

    public ScopeContainer() {

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
