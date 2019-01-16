package eu.adamwilliams.reftypes.prototype.ast;

import java.util.Optional;

public abstract class Statement {
    public abstract Optional<Expression> execute();
}
