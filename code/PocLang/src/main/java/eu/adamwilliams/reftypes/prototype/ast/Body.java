package eu.adamwilliams.reftypes.prototype.ast;

import java.util.ArrayList;
import java.util.List;

public class Body {
    private List<Statement> statements;
    private Statement parentStatement;

    public Body(List<Statement> statements, Statement parentStatement) {
        this.statements = statements;
        this.parentStatement = parentStatement;
    }

    public Body(Statement parentStatement) {
        this.statements = new ArrayList<>();
        this.parentStatement = parentStatement;
    }

    public Statement getParentStatement() {
        return parentStatement;
    }

    public List<Statement> getStatements() {
        return new ArrayList<>(this.statements);
    }
}
