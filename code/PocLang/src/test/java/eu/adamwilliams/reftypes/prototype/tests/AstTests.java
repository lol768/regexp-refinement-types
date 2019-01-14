package eu.adamwilliams.reftypes.prototype.tests;

import eu.adamwilliams.reftypes.prototype.Application;
import eu.adamwilliams.reftypes.prototype.ast.*;
import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeCheckResults;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;
import org.antlr.v4.runtime.tree.ParseTree;
import org.junit.Assert;
import org.junit.Test;

public class AstTests {
    @Test
    public void testMath() {
        BinaryOperationExpression binOpAdd = new BinaryOperationExpression(
                BinaryOperationType.INT_ADD,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(7L, (long) ((Long) binOpAdd.evaluate()));

        BinaryOperationExpression binOpSub = new BinaryOperationExpression(
                BinaryOperationType.INT_SUBTRACT,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(1L, (long) ((Long) binOpSub.evaluate()));

        BinaryOperationExpression binOpMultiply = new BinaryOperationExpression(
                BinaryOperationType.INT_MULTIPLY,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(12L, (long) ((Long) binOpMultiply.evaluate()));

        BinaryOperationExpression binOpDivide = new BinaryOperationExpression(
                BinaryOperationType.INT_DIVIDE,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(2, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(2L, (long) ((Long) binOpDivide.evaluate()));
    }

    @Test
    public void testStringConcat() {
        BinaryOperationExpression binOpConcat = new BinaryOperationExpression(
                BinaryOperationType.STRING_CONCAT,
                new StringLiteral("a", new TypeContainer(Type.STRING, null)),
                new StringLiteral("b", new TypeContainer(Type.STRING, null))
        );
        Assert.assertEquals("ab", binOpConcat.evaluate());
    }

    @Test
    public void testMathMoreComplex() {
        BinaryOperationExpression binOpSubEx = new BinaryOperationExpression(
                BinaryOperationType.INT_ADD,
                new UIntLiteral(4, new TypeContainer(Type.UNSIGNED_INTEGER, null)),
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );

        BinaryOperationExpression binOpEx = new BinaryOperationExpression(
                BinaryOperationType.INT_MULTIPLY,
                binOpSubEx,
                new UIntLiteral(3, new TypeContainer(Type.UNSIGNED_INTEGER, null))
        );
        Assert.assertEquals(21L, (long) ((Long) binOpEx.evaluate()));
    }

    @Test
    public void testBasicAstConstruction() {
        String basicProgram = "function LookupUserById(id: uint[> 1]): void {\n" +
                "    return\n" +
                "}";

        ParseTree tree = ParsingTests.getParseTree(basicProgram);
        Application app = new Application();
        TypeCheckResults results = app.doTypeChecks(tree);
        // I miss shouldly :(

        Assert.assertTrue("Function table is populated properly", results.getFunctionTable().hasFunction("LookupUserById"));
        Assert.assertEquals("Number of arguments is correct", 1, results.getFunctionTable().getFunctionByIdentifier("LookupUserById").getArguments().size());
        Assert.assertEquals("Number of body statements is correct", 1, results.getFunctionTable().getFunctionByIdentifier("LookupUserById").getBody().getStatements().size());
        Assert.assertTrue("Body statement type is correct", results.getFunctionTable().getFunctionByIdentifier("LookupUserById").getBody().getStatements().get(0) instanceof ReturnStatement);
    }

    @Test
    public void testNestedIfAst() {
        String program = "function FetchString(): string[/.+/] {\n" +
                "    if (true) {\n" +
                "        if (false) {\n" +
                "            return \"foo\"\n" +
                "        }\n" +
                "    }\n" +
                "}";
        ParseTree tree = ParsingTests.getParseTree(program);
        Application app = new Application();
        TypeCheckResults results = app.doTypeChecks(tree);
        Body body = results.getFunctionTable().getFunctionByIdentifier("FetchString").getBody();
        Assert.assertTrue("If statement for first statement of body", body.getStatements().get(0) instanceof IfStatement);

    }

    @Test
    public void testEvaluateReturn() {
        String moreAdvancedProgram = "function Test(id: uint[> 1]): uint {\n" +
                "    return 1+1\n" +
                "}";

        ParseTree tree = ParsingTests.getParseTree(moreAdvancedProgram);
        Application app = new Application();
        TypeCheckResults typeCheckResults = app.doTypeChecks(tree);
        FunctionDeclaration function = typeCheckResults.getFunctionTable().getFunctionByIdentifier("Test");

        ReturnStatement returnStmt = (ReturnStatement) function.getBody().getStatements().get(0);
        Assert.assertTrue(returnStmt.getValue().isPresent());
        long result = (long) returnStmt.getValue().get().evaluate();
        Assert.assertEquals(2, result);
    }

    @Test
    public void testVariableAssignment() {
        String moreAdvancedProgram = "function Test(id: uint[> 1]): uint {\n" +
                "    var a: uint\n" +
                "    a=4+1\n" +
                "    return a\n" +
                "}";

        ParseTree tree = ParsingTests.getParseTree(moreAdvancedProgram);
        Application app = new Application();
        TypeCheckResults typeCheckResults = app.doTypeChecks(tree);
        FunctionDeclaration function = typeCheckResults.getFunctionTable().getFunctionByIdentifier("Test");

        VariableAssignmentStatement assignStmt = (VariableAssignmentStatement) function.getBody().getStatements().get(1);
        assignStmt.execute();
        ReturnStatement returnStmt = (ReturnStatement) function.getBody().getStatements().get(2);

        Assert.assertTrue(returnStmt.getValue().isPresent());
        long result = (long) returnStmt.getValue().get().evaluate();
        Assert.assertEquals(5, result);
    }
}

// End of term 2, start of Easter break