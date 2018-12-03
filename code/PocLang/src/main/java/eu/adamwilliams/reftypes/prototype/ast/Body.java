package eu.adamwilliams.reftypes.prototype.ast;

import java.util.ArrayList;
import java.util.List;

public class Body {
    private List<Statement> statements;

    public Body(List<Statement> statements) {
        this.statements = statements;
    }

    public List<Statement> getStatements() {
        return new ArrayList<>(this.statements);
    }
}
