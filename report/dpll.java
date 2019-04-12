// Set up the symbols and literals
DPLLSatisfiable dpllSatisfiable = new DPLLSatisfiable();
PropositionSymbol a = new PropositionSymbol("A");
PropositionSymbol b = new PropositionSymbol("B");
PropositionSymbol c = new PropositionSymbol("C");

ComplexSentence notA = new ComplexSentence(Connective.NOT, a);
ComplexSentence notC = new ComplexSentence(Connective.NOT, c);

// (a \lor \neg b) \land (\neg a \lor \neg b) \land c
Sentence problem = Sentence.newConjunction(
    Sentence.newDisjunction(a, b),
    Sentence.newDisjunction(c, notA),
    Sentence.newDisjunction(notA, notC)
);

Set<Clause> clauses = ConvertToConjunctionOfClauses.convert(problem).getClauses();
List<PropositionSymbol> result = new ArrayList<>(
    SymbolCollector.getSymbolsFrom(problem)
);

// Create empty model
Model model = new Model();
// Do the satisfiability check
boolean dpllResult = dpllSatisfiable.dpll(clauses, result, model);
System.out.println(dpllResult); // true = SAT, false = UNSAT