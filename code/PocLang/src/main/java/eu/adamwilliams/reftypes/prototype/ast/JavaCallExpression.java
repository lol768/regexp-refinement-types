package eu.adamwilliams.reftypes.prototype.ast;

import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.lang.reflect.Modifier;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;
import java.util.Optional;

public class JavaCallExpression extends Expression {

    private String[] components;
    private List<Expression> arguments;


    public JavaCallExpression(String[] components, List<Expression> arguments) {
        this.components = components;
        this.arguments = arguments;
    }

    @Override
    public Object evaluate() {
        // we implement the Java "contextually ambiguous names" algorithm
        // or, at least, a subset of it (because some of the contexts don't apply to us).

        // https://docs.oracle.com/javase/specs/jls/se7/html/jls-6.html#jls-6.5
        StringBuilder joined = new StringBuilder();
        String[] rest = null;
        Class<?> theClass = null;
        for (int i = this.components.length; i >= 1; i--) {
            joined.setLength(0);
            for (int j = 0; j < i; j++) {
                joined.append(this.components[j] + ".");
            }
            try {
                String substring = joined.substring(0, joined.lastIndexOf("."));
                theClass = Class.forName(substring);
                rest = new String[this.components.length-i];
                System.arraycopy(this.components, i, rest, 0, this.components.length-i);
                break;
            } catch (ClassNotFoundException ignored) {
            }
        }

        if (theClass == null) {
            throw new IllegalArgumentException("Provided target is not a valid class");
        }

        Object instance = null;
        for (int i = 0; i < rest.length; i++) {
            String[] finalRest = rest;
            int finalI = i;

            Optional<Method> candidateMethod = Arrays.stream(theClass.getMethods()).filter(f -> f.getName().equals(finalRest[finalI]) && transformArguments(this.arguments, f.getParameterTypes()) != null).findAny();
            Optional<Field> candidateField = Arrays.stream(theClass.getFields()).filter(f -> f.getName().equals(finalRest[finalI])).findAny();
            Object[] methodArgs = candidateMethod.map(method -> transformArguments(this.arguments, method.getParameterTypes())).orElse(null);

            if (candidateField.isPresent() && (instance != null || Modifier.isStatic(candidateField.get().getModifiers()))) {
                try {
                    instance = candidateField.get().get(instance);
                    theClass = instance.getClass();
                } catch (IllegalAccessException e) {
                    e.printStackTrace();
                    return null; // what can we do here? Nothing.
                }
            } else if (candidateMethod.isPresent() && methodArgs != null && (instance != null || Modifier.isStatic(candidateMethod.get().getModifiers()))) {
                try {
                    instance = candidateMethod.get().invoke(instance, methodArgs);
                    if (instance != null) {
                        theClass = instance.getClass();
                    }
                } catch (IllegalAccessException | InvocationTargetException e) {
                    return null; // what can we do here? Nothing.
                }
            }
        }


        return instance;
    }

    public static Object[] transformArguments(List<Expression> arguments, Class<?>[] parameterTypes) {
        Object[] argumentsToPassToJavaMethod = new Object[parameterTypes.length];
        for (int i = 0; i < arguments.size(); i++) {
            Expression argument = arguments.get(i);
            Object resultAfterEvaluation = argument.evaluate();
            Class<?> correspondingType = parameterTypes[i];
            if ((resultAfterEvaluation instanceof Long && correspondingType.isAssignableFrom(Long.class)) || (resultAfterEvaluation instanceof Boolean && correspondingType.isAssignableFrom(Boolean.class)) || (resultAfterEvaluation instanceof String && correspondingType.isAssignableFrom(String.class))) {
                argumentsToPassToJavaMethod[i] = correspondingType.cast(resultAfterEvaluation);
                continue;
            }

            return null;
        }
        return argumentsToPassToJavaMethod;
    }
}
