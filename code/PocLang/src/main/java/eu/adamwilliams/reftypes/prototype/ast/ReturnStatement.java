package eu.adamwilliams.reftypes.prototype.ast;

import java.util.Optional;

public class ReturnStatement extends Statement {
    private Expression value;

    public ReturnStatement(Expression value) {
        this.value = value;
    }

    public Optional<Expression> getValue() {
        return Optional.ofNullable(value);
    }

    @Override
    public void execute() {
        // TODO
    }
}
