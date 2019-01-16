package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

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
    public Optional<Expression> execute() {
        if (this.value == null) {
            // TODO: Sort this out later..
            return Optional.of(new StringLiteral("VOID", new TypeContainer(Type.STRING, null)));
        }
        return Optional.of(this.value);
    }
}
