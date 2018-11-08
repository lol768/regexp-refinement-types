package eu.adamwilliams.reftypes.prototype;

import eu.adamwilliams.reftypes.prototype.domain.*;
import eu.adamwilliams.reftypes.prototype.domain.FunctionDeclaration;
import eu.adamwilliams.reftypes.prototype.parser.PocLangBaseListener;
import eu.adamwilliams.reftypes.prototype.parser.PocLangParser;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.Token;
import org.fusesource.jansi.Ansi;
import org.sosy_lab.java_smt.api.*;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

/**
 * Visitor for items in the parse tree. We'll do things
 * like register functions in the function table here.
 */
public class VisitorListener extends PocLangBaseListener {
    private FunctionTable table;
    private ErrorReporter reporter;
    private SolverContext context;
    private VisitorPhase phase = VisitorPhase.COLLECTING_FUNCTIONS;
    private Map<String, StackEntry> typesCurrentlyInScope = new HashMap<>();
    private IntegerFormulaManager ifm;
    private NumeralFormula.IntegerFormula x;

    public VisitorListener(FunctionTable table, ErrorReporter reporter, SolverContext context) {
        this.table = table;
        this.reporter = reporter;
        this.context = context;
        this.ifm = context.getFormulaManager().getIntegerFormulaManager();
        this.x = ifm.makeVariable("x");
    }

    public void setPhase(VisitorPhase phase) {
        this.phase = phase;
    }


    @Override
    public void enterFunction_sig(PocLangParser.Function_sigContext ctx) {
        String idText = ctx.IDENTIFIER().getText();
        Token symbol = ctx.IDENTIFIER().getSymbol();

        if (this.phase == VisitorPhase.COLLECTING_FUNCTIONS) {
            if (this.table.hasFunction(idText)) {
                FunctionDeclaration function = this.table.getFunctionByIdentifier(idText);
                String prevLocation = String.format("L%d:%d", function.getLineNo(), function.getColNo());
                String msg = String.format("Redeclaration of function %s, previously declared at %s", ctx.IDENTIFIER(), prevLocation);
                this.reporter.reportError(new ErrorReport(symbol, msg));
                return;
            }

            LinkedHashMap<String, TypeContainer> collect = ctx.argument_decl().stream().collect(
                    Collectors.toMap(
                            ad -> ad.IDENTIFIER().getText(),
                            ad -> mapTypeFromParsed(ad.type()),
                            (u, v) -> {
                                throw new IllegalStateException(String.format("Duplicate key %s", u));
                            },
                            LinkedHashMap::new
                    )
            );
            this.table.addFunction(new FunctionDeclaration(idText, mapTypeFromParsed(ctx.type()), collect, symbol.getLine(), symbol.getCharPositionInLine()));

        } else if (this.phase == VisitorPhase.CHECKING_TYPES) {
            LinkedHashMap<String, TypeContainer> arguments = this.table.getFunctionByIdentifier(idText).getArguments();
            for (Map.Entry<String, TypeContainer> entry : arguments.entrySet()) {
                this.typesCurrentlyInScope.put(entry.getKey(), new StackEntry(entry.getValue(), StackEntryType.ARGUMENT));
            }
        }
    }

    private TypeContainer mapTypeFromParsed(PocLangParser.TypeContext ctx) {
        PocLangParser.Type_keywordContext keyword = ctx.type_keyword();
        Type type;
        if (keyword instanceof PocLangParser.StringTypeContext) {
            type = Type.STRING;
        } else if (keyword instanceof PocLangParser.VoidTypeContext) {
            type = Type.VOID;
        } else if (keyword instanceof PocLangParser.UnsignedIntTypeContext) {
            type = Type.UNSIGNED_INTEGER;
        } else {
            throw new IllegalArgumentException("Unexpected type " + keyword.getClass().getName());
        }

        if (!this.ensureApplicableConstraint(ctx, type)) {
            // we just ignore the constraint and continue, the error is logged
            return new TypeContainer(type, null);
        }

        BooleanFormula bf = null;

        // we use the inverse in our boolean formulae
        if (ctx.int_constraint() != null) {
            if (ctx.int_constraint() instanceof PocLangParser.GreaterThanConstraintContext) {
                String val = ((PocLangParser.GreaterThanConstraintContext) ctx.int_constraint()).INT().getText();
                bf = ifm.lessOrEquals(x, ifm.makeNumber(val));
            }
            if (ctx.int_constraint() instanceof PocLangParser.LessThanConstraintContext) {
                String val = ((PocLangParser.LessThanConstraintContext) ctx.int_constraint()).INT().getText();
                bf = ifm.greaterOrEquals(x, ifm.makeNumber(val));
            }
            if (ctx.int_constraint() instanceof PocLangParser.GreaterThanEqualsConstraintContext) {
                String val = ((PocLangParser.GreaterThanEqualsConstraintContext) ctx.int_constraint()).INT().getText();
                bf = ifm.lessThan(x, ifm.makeNumber(val));
            }
            if (ctx.int_constraint() instanceof PocLangParser.LessThanEqualsConstraintContext) {
                String val = ((PocLangParser.LessThanEqualsConstraintContext) ctx.int_constraint()).INT().getText();
                bf = ifm.greaterThan(x, ifm.makeNumber(val));
            }
        }
        return new TypeContainer(type, bf);
    }

    private boolean ensureApplicableConstraint(PocLangParser.TypeContext ctx, Type type) {
        switch(type) {
            case UNSIGNED_INTEGER:
                if (ctx.string_constraint() != null) {
                    this.reporter.reportError(new ErrorReport(ctx.type_keyword().start, "Attempt to apply string constraint to uint type."));
                    return false;
                }
                break;
            case STRING:
                if (ctx.int_constraint() != null) {
                    this.reporter.reportError(new ErrorReport(ctx.type_keyword().start, "Attempt to apply int constraint to string type."));
                    return false;
                }
                break;
            case VOID:
                break;
        }
        return true;
    }

    private TypeContainer inferTypeFromValue(PocLangParser.Value_refContext value_refContext) {
        if (value_refContext.identifier_ref() == null) {
            Type type = value_refContext.INT() != null ? Type.UNSIGNED_INTEGER : Type.STRING; // FIXME: Do this better

            if (value_refContext.INT() != null) {
                return new TypeContainer(type, ifm.equal(x, ifm.makeNumber(value_refContext.INT().getText())));
            }
            return new TypeContainer(type, null);
        }

        String variableId = value_refContext.identifier_ref().getText();
        if (!this.typesCurrentlyInScope.containsKey(variableId)) {
            throw new IllegalArgumentException("Reference to variable " + variableId + " which isn't in scope");
        }
        return this.typesCurrentlyInScope.get(variableId).getType();
    }


    @Override
    public void enterVar_decl(PocLangParser.Var_declContext ctx) {
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.typesCurrentlyInScope.put(ctx.IDENTIFIER().getText(), new StackEntry(mapTypeFromParsed(ctx.type()), StackEntryType.LOCAL));
        }
    }

    @Override
    public void enterFunction_call(PocLangParser.Function_callContext ctx) {
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            String idText = ctx.IDENTIFIER().getText();
            Token symbol = ctx.IDENTIFIER().getSymbol();

            if (!this.table.hasFunction(idText)) {
                String msg = String.format("Call to undefined function with identifier %s", idText);
                this.reporter.reportError(new ErrorReport(symbol, msg));
                return;
            }

            // check types of arguments
            for (int i = 0; i < ctx.expr().size(); i++) {
                PocLangParser.ExprContext expr = ctx.expr(i);
                checkExprType(expr, this.table.getFunctionByIdentifier(idText).getNthArgument(i).getValue(), symbol, reporter);
            }
        }
    }

    private void checkExprType(PocLangParser.ExprContext expr, TypeContainer expected, Token symbol, ErrorReporter reporter) {
        PocLangParser.Value_refContext value_refContext = expr.value_ref();
        if (value_refContext != null) {
            try {
                TypeContainer tc = this.inferTypeFromValue(value_refContext);
                if (!this.checkTypes(expected, tc)) {
                    reporter.reportError(new ErrorReport(symbol, tc + " didn't satisfy " + expected));
                }
            } catch (IllegalArgumentException ex) {
                reporter.reportError(new ErrorReport(symbol, ex.getMessage()));
            }
        }

        PocLangParser.Function_callContext function_callContext = expr.function_call();
        if (function_callContext != null && this.table.hasFunction(function_callContext.IDENTIFIER().getText())) {
            FunctionDeclaration functionByIdentifier = this.table.getFunctionByIdentifier(function_callContext.IDENTIFIER().getText());
            if (!this.checkTypes(expected, functionByIdentifier.getReturnType())) {
                reporter.reportError(new ErrorReport(symbol, "Return type " + functionByIdentifier.getReturnType() + " of function " + functionByIdentifier.getIdentifier() + " didn't satisfy " + expected));
            }
        }
    }

    @Override
    public void enterVar_assignment(PocLangParser.Var_assignmentContext ctx) {
        String variableId = ctx.IDENTIFIER().getText();

        if (!this.typesCurrentlyInScope.containsKey(variableId)) {
            this.reporter.reportError(new ErrorReport(ctx.getStart(), "Assignment to variable " + variableId + " which isn't in scope"));
            return; // bail, we can't check types at this point
        }
        this.checkExprType(ctx.expr(), this.typesCurrentlyInScope.get(variableId).getType(), ctx.getStart(), this.reporter);
    }

    private boolean checkTypes(TypeContainer expected, TypeContainer actual) {
        if (expected.getType() != actual.getType()) {
            return false; // fail fast
        }

        if (expected.getRefinement() == null) {
            return true;
        }

        // Solve formula, get model, and print variable assignment
        try (ProverEnvironment prover = context.newProverEnvironment(SolverContext.ProverOptions.GENERATE_MODELS)) {
            if (actual.getRefinement() != null) {
                prover.addConstraint(actual.getRefinement());
            }

            prover.addConstraint(expected.getRefinement());
            boolean isUnsat = prover.isUnsat();
            if (isUnsat) {
                return true;
            }
            try (Model model = prover.getModel()) {
                System.out.printf(Ansi.ansi().fg(Ansi.Color.RED)+"// Solved SAT with  = %s%n", model.evaluate(x));
            }
        } catch (InterruptedException | SolverException e) {
            e.printStackTrace();
            System.exit(-1); // can't do anything useful here
        }
        return false;
    }

    @Override
    public void enterReturn_stmt(PocLangParser.Return_stmtContext ctx) {
        // this is horribly ugly, but we need to grab the enclosing function somehow
        ParserRuleContext bodyLineCtx = ctx.getParent();
        PocLangParser.BodyContext bodyCtx = (PocLangParser.BodyContext) bodyLineCtx.getParent();
        PocLangParser.FunctionContext function = (PocLangParser.FunctionContext) bodyCtx.getParent();
        PocLangParser.Function_sigContext functionSignature = function.function_sig();
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            FunctionDeclaration decl = this.table.getFunctionByIdentifier(functionSignature.IDENTIFIER().getText());
            if (ctx.expr() == null && decl.getReturnType().getType() == Type.VOID) {
                return;
            } else if (decl.getReturnType().getType() == Type.VOID) {
                this.reporter.reportError(new ErrorReport(ctx.expr().start, "Attempt to return value from void function " + functionSignature.IDENTIFIER().getText()));
            }
            this.checkExprType(ctx.expr(), decl.getReturnType(), ctx.getStart(), this.reporter);
        }
    }

    @Override
    public void exitBody(PocLangParser.BodyContext ctx) {
        List<PocLangParser.Body_lineContext> lines = ctx.body_line();
        boolean hasAtLeastOneReturn = lines.stream().flatMap(bl -> bl.children.stream()).anyMatch(line -> line instanceof PocLangParser.Return_stmtContext);
        PocLangParser.Function_sigContext functionSignature = (PocLangParser.Function_sigContext) ctx.getParent().children.get(0);
        FunctionDeclaration functionDeclaration = this.table.getFunctionByIdentifier(functionSignature.IDENTIFIER().getText());

        if (!hasAtLeastOneReturn && this.phase == VisitorPhase.COLLECTING_FUNCTIONS && functionDeclaration.getReturnType().getType() != Type.VOID) {
            String owningFunctionId = functionDeclaration.getIdentifier();
            this.reporter.reportError(new ErrorReport(lines.get(lines.size() - 1).getStart(), "No return statement in function " + owningFunctionId));
            return;
        }
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.typesCurrentlyInScope.clear();
        }
    }


}