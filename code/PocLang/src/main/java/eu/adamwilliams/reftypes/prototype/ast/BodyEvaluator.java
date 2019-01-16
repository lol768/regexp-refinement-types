package eu.adamwilliams.reftypes.prototype.ast;

import java.util.Optional;

public class BodyEvaluator {

    public static Optional<Expression> evaluateBody(Body b) {
        int i = 0;
        Statement stmt = b.getStatements().get(i);
        Optional<Expression> result = null;
        while (!(result = stmt.execute()).isPresent()) {
            stmt = b.getStatements().get(i++);
        }

        return result;
    }
}
