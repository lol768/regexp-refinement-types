IntegerFormulaManager imgr = ctx.getFormulaManager().getIntegerFormulaManager();
IntegerFormula x = imgr.makeVariable("x");

BooleanFormula expectedNegated = imgr.lessOrEquals(x, x.makeNumber(4)); // x <= 4
BooleanFormula actual = imgr.lessThan(x, x.makeNumber(10)); // x < 10

// try-with-resources, will dispose of the ProverEnvironment for us
try (ProverEnvironment prover = ctx.newProverEnvironment(ProverOptions.GENERATE_MODELS)) {
    prover.addConstraint(expectedNegated);
    prover.addConstraint(actual); // both constraints must hold
    boolean isUnsat = prover.isUnsat(); // isUnsat = false here.
    try (Model model = prover.getModel()) {
        // Prints the value of x found for which the two constraints hold
        // (i.e. a violation of the refinement type has been found)
        System.out.printf("Violation found with: ", model.evaluate(x));
    }
}