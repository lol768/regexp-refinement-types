@Test
public void testJavaCall() {
    // Arrange program
    String moreAdvancedProgram = "function Main(): uint {\n" +
            "    var a: uint\n" +
            "    a=!java.lang.Long.divideUnsigned(10,2)\n" +
            "    return a\n" +
            "}";

    // Act - parse the program
    ParseTree tree = ParsingTests.getParseTree(moreAdvancedProgram);

    // Set up and invoke the type checker
    Application app = new Application();
    TypeCheckResults results = app.doTypeChecks(tree);

    // Assert - the type checker should not raise any errors
    Assert.assertEquals(
            "No error reports should have been raised",
            0, // expected number
            results.getReports().size() // actual
    );

    // Point the interpreter at the Main function and execute
    FunctionTable table = results.getFunctionTable();
    FunctionDeclaration main = table.getFunctionByIdentifier("Main");
    Optional<Expression> result = BodyEvaluator.evaluateBody(main.getBody());

    // Assert - ensure that the result of the execution is 5
    Assert.assertEquals(5L, result.get().evaluate());
}