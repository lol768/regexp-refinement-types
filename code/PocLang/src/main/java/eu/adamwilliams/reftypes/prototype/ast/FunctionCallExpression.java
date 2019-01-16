package eu.adamwilliams.reftypes.prototype.ast;

import java.util.List;

public class FunctionCallExpression extends Expression {

    private FunctionDeclaration callee;
    private List<Expression> arguments;

    public FunctionCallExpression(FunctionDeclaration callee, List<Expression> arguments) {
        this.callee = callee;
        this.arguments = arguments;
    }

    public FunctionDeclaration getCallee() {
        return callee;
    }

    public List<Expression> getArguments() {
        return arguments;
    }

    @Override
    public Object evaluate() {
        for (int i = 0; i < this.arguments.size(); i++) {
            callee.getArgumentEntries().get(i).setCurrentValue(this.arguments.get(i).evaluate());
        }
        Body calleeBody = callee.getBody();
        return BodyEvaluator.evaluateBody(calleeBody).get().evaluate();
    }
}
