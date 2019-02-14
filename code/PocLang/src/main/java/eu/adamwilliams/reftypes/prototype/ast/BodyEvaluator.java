package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.StackEntry;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

import java.util.List;
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

    public static Optional<Expression> evaluateFunction(FunctionDeclaration fd, Object... args) {
        List<StackEntry> argumentEntries = fd.getArgumentEntries();
        if (args.length != argumentEntries.size()) {
            throw new IllegalArgumentException("Number of arguments passed must equal number required by function");
        }

        for (int i = 0; i < args.length; i++) {
            TypeContainer type = argumentEntries.get(i).getType();
            TypeContainer actual = JavaCallExpression.getTypeContainerFromJavaObject(args[i].getClass());
            if (actual == null || type.getType() != actual.getType()) {
                throw new IllegalArgumentException("Argument " + i + " type is invalid, needs to be assignable to " + type.getType());
            }
            argumentEntries.get(i).setCurrentValue(args[i]);
        }
        return evaluateBody(fd.getBody());
    }
}
