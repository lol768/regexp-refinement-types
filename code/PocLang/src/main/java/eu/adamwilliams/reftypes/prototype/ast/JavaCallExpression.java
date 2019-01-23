package eu.adamwilliams.reftypes.prototype.ast;

import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;

import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.lang.reflect.Modifier;
import java.util.Arrays;
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
    public TypeContainer getTypeContainer() {
        Class<?> javaType = this.getTypeOfResult();
        if (javaType.isAssignableFrom(Long.class) || javaType == Long.TYPE) {
            return new TypeContainer(Type.UNSIGNED_INTEGER, null);
        } else if (javaType.isAssignableFrom(Boolean.class) || javaType == Boolean.TYPE) {
            return new TypeContainer(Type.BOOLEAN, null);
        } else if (javaType.isAssignableFrom(String.class)) {
            return new TypeContainer(Type.STRING, null);
        } else if (javaType.isAssignableFrom(Void.class)) {
            return new TypeContainer(Type.VOID, null);
        }
        return null;
    }

    public Class<?> getTypeOfResult() {
        // TODO: dupe'd code
        String[] components = this.components;
        ClassResolutionResult classResolutionResult = new ClassResolutionResult(components).invoke();
        String[] rest = classResolutionResult.getRest();
        Class<?> theClass = classResolutionResult.getTheClass();

        if (theClass == null) {
            throw new IllegalArgumentException("Provided target is not a valid class");
        }

        boolean instance = false;
        for (int i = 0; i < rest.length; i++) {
            int finalI = i;

            Optional<Method> candidateMethod = Arrays.stream(theClass.getMethods()).filter(f -> f.getName().equals(rest[finalI]) && transformArguments(this.arguments, f.getParameterTypes()) != null).findAny();
            Optional<Field> candidateField = Arrays.stream(theClass.getFields()).filter(f -> f.getName().equals(rest[finalI])).findAny();
            Object[] methodArgs = candidateMethod.map(method -> transformArguments(this.arguments, method.getParameterTypes())).orElse(null);

            if (candidateField.isPresent() && (instance || Modifier.isStatic(candidateField.get().getModifiers()))) {
                instance = true;
                theClass = candidateField.get().getType();
            } else if (candidateMethod.isPresent() && methodArgs != null && (instance || Modifier.isStatic(candidateMethod.get().getModifiers()))) {
                theClass = candidateMethod.get().getReturnType();
            }
        }
        return theClass;
    }

    @Override
    public Object evaluate() {
        // we implement the Java "contextually ambiguous names" algorithm
        // or, at least, a subset of it (because some of the contexts don't apply to us).

        // https://docs.oracle.com/javase/specs/jls/se7/html/jls-6.html#jls-6.5
        String[] components = this.components;
        ClassResolutionResult classResolutionResult = new ClassResolutionResult(components).invoke();
        String[] rest = classResolutionResult.getRest();
        Class<?> theClass = classResolutionResult.getTheClass();

        if (theClass == null) {
            throw new IllegalArgumentException("Provided target is not a valid class");
        }

        Object instance = null;
        for (int i = 0; i < rest.length; i++) {
            int finalI = i;

            Optional<Method> candidateMethod = Arrays.stream(theClass.getMethods()).filter(f -> f.getName().equals(rest[finalI]) && transformArguments(this.arguments, f.getParameterTypes()) != null).findAny();
            Optional<Field> candidateField = Arrays.stream(theClass.getFields()).filter(f -> f.getName().equals(rest[finalI])).findAny();
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
            if ((resultAfterEvaluation instanceof Long && (correspondingType.isAssignableFrom(Long.class))) || (resultAfterEvaluation instanceof Boolean && correspondingType.isAssignableFrom(Boolean.class)) || (resultAfterEvaluation instanceof String && correspondingType.isAssignableFrom(String.class))) {
                argumentsToPassToJavaMethod[i] = correspondingType.cast(resultAfterEvaluation);
                continue;
            }

            if (resultAfterEvaluation instanceof Long && (correspondingType == Long.TYPE) || resultAfterEvaluation instanceof Boolean && (correspondingType == Boolean.TYPE)) {
                argumentsToPassToJavaMethod[i] = resultAfterEvaluation; // primitive.
                continue;
            }

            return null;
        }
        return argumentsToPassToJavaMethod;
    }

    private class ClassResolutionResult {
        private String[] components;
        private String[] rest;
        private Class<?> theClass;

        public ClassResolutionResult(String... components) {
            this.components = components;
        }

        public String[] getRest() {
            return rest;
        }

        public Class<?> getTheClass() {
            return theClass;
        }

        public ClassResolutionResult invoke() {
            StringBuilder joined = new StringBuilder();
            rest = null;
            theClass = null;
            for (int i = components.length; i >= 1; i--) {
                joined.setLength(0);
                for (int j = 0; j < i; j++) {
                    joined.append(components[j] + ".");
                }
                try {
                    String substring = joined.substring(0, joined.lastIndexOf("."));
                    theClass = Class.forName(substring);
                    rest = new String[components.length - i];
                    System.arraycopy(components, i, rest, 0, components.length - i);
                    break;
                } catch (ClassNotFoundException ignored) {
                }
            }
            return this;
        }
    }
}
