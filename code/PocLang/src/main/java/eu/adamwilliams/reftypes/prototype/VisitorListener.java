package eu.adamwilliams.reftypes.prototype;

import com.microsoft.z3.*;
import eu.adamwilliams.reftypes.prototype.domain.*;
import eu.adamwilliams.reftypes.prototype.domain.FunctionDeclaration;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import eu.adamwilliams.reftypes.prototype.parser.PocLangBaseListener;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.Token;

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
    private final RegexZ3Adapter regexAdapter;
    private FunctionTable table;
    private ErrorReporter reporter;
    private final Context z3Ctx;
    private VisitorPhase phase = VisitorPhase.COLLECTING_FUNCTIONS;
    private Map<String, StackEntry> typesCurrentlyInScope = new HashMap<>();
    private IntExpr x; // represents unknown int
    private SeqExpr y; // represents unknown string

    public VisitorListener(FunctionTable table, ErrorReporter reporter, Context z3Ctx) {
        this.table = table;
        this.reporter = reporter;
        this.z3Ctx = z3Ctx;
        this.x = z3Ctx.mkIntConst("x");
        this.y = (SeqExpr) z3Ctx.mkConst(z3Ctx.mkSymbol("x"), z3Ctx.getStringSort());
        this.regexAdapter = new RegexZ3Adapter();
    }

    public void setPhase(VisitorPhase phase) {
        this.phase = phase;
    }


    @Override
    public void enterFunction_sig(PocLang.Function_sigContext ctx) {
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

    private TypeContainer mapTypeFromParsed(PocLang.TypeContext ctx) {
        PocLang.Type_keywordContext keyword = ctx.type_keyword();
        Type type;
        if (keyword instanceof PocLang.StringTypeContext) {
            type = Type.STRING;
        } else if (keyword instanceof PocLang.VoidTypeContext) {
            type = Type.VOID;
        } else if (keyword instanceof PocLang.UnsignedIntTypeContext) {
            type = Type.UNSIGNED_INTEGER;
        } else {
            throw new IllegalArgumentException("Unexpected type " + keyword.getClass().getName());
        }

        if (!this.ensureApplicableConstraint(ctx, type)) {
            // we just ignore the constraint and continue, the error is logged
            return new TypeContainer(type, null);
        }

        BoolExpr bf = null;

        if (ctx.int_constraint() != null) {
            if (ctx.int_constraint() instanceof PocLang.GreaterThanConstraintContext) {
                String val = ((PocLang.GreaterThanConstraintContext) ctx.int_constraint()).CONSTRAINT_UINT().getText();
                bf = z3Ctx.mkGt(x, z3Ctx.mkInt(val));
            }
            if (ctx.int_constraint() instanceof PocLang.LessThanConstraintContext) {
                String val = ((PocLang.LessThanConstraintContext) ctx.int_constraint()).CONSTRAINT_UINT().getText();
                bf = z3Ctx.mkLt(x, z3Ctx.mkInt(val));
            }
            if (ctx.int_constraint() instanceof PocLang.GreaterThanEqualsConstraintContext) {
                String val = ((PocLang.GreaterThanEqualsConstraintContext) ctx.int_constraint()).CONSTRAINT_UINT().getText();
                bf = z3Ctx.mkGe(x, z3Ctx.mkInt(val));
            }
            if (ctx.int_constraint() instanceof PocLang.LessThanEqualsConstraintContext) {
                String val = ((PocLang.LessThanEqualsConstraintContext) ctx.int_constraint()).CONSTRAINT_UINT().getText();
                bf = z3Ctx.mkLe(x, z3Ctx.mkInt(val));
            }
        } else if (ctx.string_constraint() != null) {
            bf = this.regexAdapter.getRefinementType(this.y, ctx.string_constraint().re(), this.z3Ctx);
        }
        if (bf == null) {
            // type is unconstrained
            return new TypeContainer(type, null);
        }

        bf = z3Ctx.mkNot(bf); // we want to find violations
        return new TypeContainer(type, bf);
    }

    private boolean ensureApplicableConstraint(PocLang.TypeContext ctx, Type type) {
        switch (type) {
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

    private TypeContainer inferTypeFromValue(PocLang.Value_refContext value_refContext) {
        if (value_refContext.identifier_ref() == null) {
            Type type = value_refContext.INT() != null ? Type.UNSIGNED_INTEGER : Type.STRING; // FIXME: Do this better

            if (value_refContext.INT() != null) {
                BoolExpr eqExpr = z3Ctx.mkEq(x, z3Ctx.mkInt(value_refContext.INT().getText()));
                return new TypeContainer(type, eqExpr);
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
    public void enterVar_decl(PocLang.Var_declContext ctx) {
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.typesCurrentlyInScope.put(ctx.IDENTIFIER().getText(), new StackEntry(mapTypeFromParsed(ctx.type()), StackEntryType.LOCAL));
        }
    }

    @Override
    public void enterFunction_call(PocLang.Function_callContext ctx) {
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
                PocLang.ExprContext expr = ctx.expr(i);
                checkExprType(expr, this.table.getFunctionByIdentifier(idText).getNthArgument(i).getValue(), symbol, reporter);
            }
        }
    }

    private void checkExprType(PocLang.ExprContext expr, TypeContainer expected, Token symbol, ErrorReporter reporter) {
        PocLang.Value_refContext value_refContext = expr.value_ref();
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

        PocLang.Function_callContext function_callContext = expr.function_call();
        if (function_callContext != null && this.table.hasFunction(function_callContext.IDENTIFIER().getText())) {
            FunctionDeclaration functionByIdentifier = this.table.getFunctionByIdentifier(function_callContext.IDENTIFIER().getText());
            if (!this.checkTypes(expected, functionByIdentifier.getReturnType())) {
                reporter.reportError(new ErrorReport(symbol, "Return type " + functionByIdentifier.getReturnType() + " of function " + functionByIdentifier.getIdentifier() + " didn't satisfy " + expected));
            }
        }
    }

    @Override
    public void enterVar_assignment(PocLang.Var_assignmentContext ctx) {
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
        Solver solver = z3Ctx.mkSolver();
        if (actual.getRefinement() != null) {
            solver.add(actual.getRefinement());
        }

        solver.add(expected.getRefinement());
        Status res = solver.check();
        if (res == Status.SATISFIABLE) {
            // if we're here, we have a model which violates the refinement types
            System.err.println("Violation via value " + solver.getModel().getConstInterp(actual.getType() == Type.STRING ? y : x));
            return false;
        } else if (res == Status.UNSATISFIABLE) {
            return true;
        }

        return false;
    }

    @Override
    public void enterReturn_stmt(PocLang.Return_stmtContext ctx) {
        // this is horribly ugly, but we need to grab the enclosing function somehow
        ParserRuleContext bodyLineCtx = ctx.getParent();
        PocLang.BodyContext bodyCtx = (PocLang.BodyContext) bodyLineCtx.getParent();
        PocLang.FunctionContext function = (PocLang.FunctionContext) bodyCtx.getParent();
        PocLang.Function_sigContext functionSignature = function.function_sig();
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
    public void exitBody(PocLang.BodyContext ctx) {
        List<PocLang.Body_lineContext> lines = ctx.body_line();
        boolean hasAtLeastOneReturn = lines.stream().flatMap(bl -> bl.children.stream()).anyMatch(line -> line instanceof PocLang.Return_stmtContext);
        PocLang.Function_sigContext functionSignature = (PocLang.Function_sigContext) ctx.getParent().children.get(0);
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