package eu.adamwilliams.reftypes.prototype;

import com.microsoft.z3.*;
import eu.adamwilliams.reftypes.prototype.ast.*;
import eu.adamwilliams.reftypes.prototype.domain.*;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import eu.adamwilliams.reftypes.prototype.parser.PocLangBaseListener;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.misc.Pair;
import org.antlr.v4.runtime.tree.ParseTree;

import java.util.*;
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
    private List<ScopeContainer> scopeContainers = new ArrayList<>();
    private IntExpr x; // represents unknown int
    private SeqExpr y; // represents unknown string
    private TypeContainer expectedIfExprType;
    private FunctionDeclaration currentFunction;
    private Statement currentStmt;

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
                            ad -> mapTypeFromParsed(ad.type_specifier()),
                            (u, v) -> {
                                throw new IllegalStateException(String.format("Duplicate key %s", u));
                            },
                            LinkedHashMap::new
                    )
            );
            this.table.addFunction(new FunctionDeclaration(idText, mapTypeFromParsed(ctx.type_specifier()), collect, symbol.getLine(), symbol.getCharPositionInLine()));
        } else if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.currentFunction = this.table.getFunctionByIdentifier(idText);
            LinkedHashMap<String, TypeContainer> arguments = this.currentFunction.getArguments();
            ScopeContainer sc = new ScopeContainer(this.currentFunction.getBody());

            for (Map.Entry<String, TypeContainer> entry : arguments.entrySet()) {
                sc.insertIdentifier(entry.getKey(), new StackEntry(entry.getValue(), StackEntryType.ARGUMENT));
            }
            this.scopeContainers.add(sc);
        }
    }

    private TypeContainer mapTypeFromParsed(PocLang.Type_specifierContext ctx) {
        PocLang.Type_keywordContext keyword = ctx.type_keyword();
        Type type;
        String friendlyRefinement = null;
        if (keyword instanceof PocLang.StringTypeContext) {
            type = Type.STRING;
        } else if (keyword instanceof PocLang.VoidTypeContext) {
            type = Type.VOID;
        } else if (keyword instanceof PocLang.UnsignedIntTypeContext) {
            type = Type.UNSIGNED_INTEGER;
        } else if (keyword instanceof PocLang.BoolTypeContext) {
            type = Type.BOOLEAN;
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
            friendlyRefinement = ctx.int_constraint().getText();
        } else if (ctx.string_constraint() != null) {
            friendlyRefinement = ctx.string_constraint().getText();
            bf = this.regexAdapter.getRefinementType(this.y, ctx.string_constraint().re(), this.z3Ctx);
        }
        if (bf == null) {
            // type is unconstrained
            return new TypeContainer(type, null);
        }

        TypeContainer typeContainer = new TypeContainer(type, bf);
        typeContainer.setFriendlyRefinement(friendlyRefinement);
        return typeContainer;
    }

    private boolean ensureApplicableConstraint(PocLang.Type_specifierContext ctx, Type type) {
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
            if (value_refContext.TRUE_LIT() != null || value_refContext.FALSE_LIT() != null) {
                type = Type.BOOLEAN;
            }
            if (value_refContext.INT() != null) {
                BoolExpr eqExpr = z3Ctx.mkEq(x, z3Ctx.mkInt(value_refContext.INT().getText()));
                return new TypeContainer(type, eqExpr);
            } else if (value_refContext.STRING_LITERAL() != null) {
                String text = value_refContext.STRING_LITERAL().getText().substring(1, value_refContext.STRING_LITERAL().getText().length() - 1);
                TypeContainer typeContainer = new TypeContainer(type, z3Ctx.mkInRe(this.y, z3Ctx.mkToRe(z3Ctx.mkString(text))));
                typeContainer.setFriendlyRefinement("\"" + text + "\"");
                return typeContainer;
            }
            return new TypeContainer(type, null);
        }

        String variableId = value_refContext.identifier_ref().getText();
        StackEntry identifierLookup = this.lookupIdentifierInScopes(variableId);

        if (identifierLookup == null) {
            this.reporter.reportError(new ErrorReport(value_refContext.start, "Couldn't infer type here," + variableId + " not in scope"));
            return new TypeContainer(Type.UNKNOWN, null);
        }
        return identifierLookup.getType();
    }


    @Override
    public void enterVar_decl(PocLang.Var_declContext ctx) {
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            String id = ctx.IDENTIFIER().getText();
            if (this.lookupIdentifierInScopes(id) != null) {
                this.reporter.reportError(new ErrorReport(
                        ctx.IDENTIFIER().getSymbol(),
                        "Declaration of '" + id + "' would shadow existing variable in scope"
                ));
            }
            this.scopeContainers.get(this.scopeContainers.size() - 1).insertIdentifier(id, new StackEntry(mapTypeFromParsed(ctx.type_specifier()), StackEntryType.LOCAL));
            this.currentFunction.getBody().getStatements().add(new VariableDeclarationStatement(id));
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
            if (ctx.getParent() instanceof PocLang.Body_lineContext) {
                FunctionCallExpression expr = getAstExpressionForFunctionCall(ctx);
                this.currentFunction.getBody().getStatements().add(new FunctionCallStatement(expr));
            }
        }
    }

    private void checkExprType(PocLang.ExprContext expr, TypeContainer expected, Token symbol, ErrorReporter reporter) {
        if (getExprType(expr) == null) {
            reporter.reportError(new ErrorReport(symbol, "Couldn't derive a type for this expression, it's probably nonsensical."));
            return;
        }

        PocLang.Value_refContext value_refContext = expr.value_ref();
        if (value_refContext != null) {
            try {
                TypeContainer tc = this.getExprType(expr);
                Pair<Boolean, String> typeCheckResult = this.checkTypes(expected, tc);
                if (!typeCheckResult.a) {
                    String supp = typeCheckResult.b != null ? ", example violating value: " + typeCheckResult.b : "";
                    reporter.reportError(new ErrorReport(symbol, tc + " didn't satisfy " + expected + supp));
                }
            } catch (IllegalArgumentException ex) {
                reporter.reportError(new ErrorReport(symbol, ex.getMessage()));
            }
        }


        PocLang.Function_callContext function_callContext = expr.function_call();
        if (function_callContext != null && this.table.hasFunction(function_callContext.IDENTIFIER().getText())) {
            FunctionDeclaration functionByIdentifier = this.table.getFunctionByIdentifier(function_callContext.IDENTIFIER().getText());
            Pair<Boolean, String> typeCheckResult = this.checkTypes(expected, getExprType(expr));
            if (!typeCheckResult.a) {
                String supp = typeCheckResult.b != null ? ", example violating value: " + typeCheckResult.b : "";
                reporter.reportError(new ErrorReport(symbol, "Return type " + functionByIdentifier.getReturnType() + " of function " + functionByIdentifier.getIdentifier() + " didn't satisfy " + expected + supp));
            }
        }
    }

    private TypeContainer getExprType(PocLang.ExprContext expr) {
        PocLang.Value_refContext value_refContext = expr.value_ref();

        if (value_refContext != null) {
            return this.inferTypeFromValue(value_refContext);
        }

        PocLang.Function_callContext function_callContext = expr.function_call();
        if (function_callContext != null && this.table.hasFunction(function_callContext.IDENTIFIER().getText())) {
            return this.table.getFunctionByIdentifier(function_callContext.IDENTIFIER().getText()).getReturnType();
        }

        List<PocLang.ExprContext> binOpItems = expr.expr();
        if (binOpItems.stream().allMatch(e -> getExprType(e) != null && getExprType(e).getType() == Type.UNSIGNED_INTEGER)) {
            return new TypeContainer(Type.UNSIGNED_INTEGER, null); // TODO: We can infer a better refinement here..
        }

        if (expr.ADD() != null && binOpItems.stream().allMatch(e -> getExprType(e) != null && getExprType(e).getType() == Type.STRING)) {
            return new TypeContainer(Type.STRING, null); // TODO: We can infer a better refinement here..
        }

        return null;
    }

    @Override
    public void enterVar_assignment(PocLang.Var_assignmentContext ctx) {
        String variableId = ctx.IDENTIFIER().getText();
        StackEntry identifierLookup = this.lookupIdentifierInScopes(variableId);
        if (identifierLookup == null) {
            this.reporter.reportError(new ErrorReport(ctx.getStart(), "Assignment to variable " + variableId + " which isn't in scope"));
            return; // bail, we can't check types at this point
        }
        this.checkExprType(ctx.expr(), identifierLookup.getType(), ctx.getStart(), this.reporter);
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.currentFunction.getBody().getStatements().add(new VariableAssignmentStatement(this.getAstExpression(ctx.expr()), variableId, identifierLookup));
        }
    }

    private Pair<Boolean, String> checkTypes(TypeContainer expected, TypeContainer actual) {
        if (expected.getType() != actual.getType()) {
            return new Pair<>(false, null); // fail fast
        }

        if (expected.getRefinement() == null) {
            return new Pair<>(true, null);
        }

        // Solve formula, get model, and print variable assignment
        Solver solver = z3Ctx.mkSolver();
        if (actual.getRefinement() != null) {
            solver.add(actual.getRefinement());
        }

        solver.add(z3Ctx.mkNot(expected.getRefinement()));
        Status res = solver.check();
        if (res == Status.SATISFIABLE) {
            // if we're here, we have a model which violates the refinement types
            return new Pair<>(false, solver.getModel().getConstInterp(actual.getType() == Type.STRING ? y : x).toString());
        } else if (res == Status.UNSATISFIABLE) {
            return new Pair<>(true, null);
        }

        return new Pair<>(false, null);
    }

    @Override
    public void enterReturn_stmt(PocLang.Return_stmtContext ctx) {
        // this is horribly ugly, but we need to grab the enclosing function somehow
        ParserRuleContext parentCtx = ctx.getParent();
        while (!(parentCtx instanceof PocLang.FunctionContext)) parentCtx = parentCtx.getParent();
        PocLang.FunctionContext function = (PocLang.FunctionContext) parentCtx;
        PocLang.Function_sigContext functionSignature = function.function_sig();
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            FunctionDeclaration decl = this.table.getFunctionByIdentifier(functionSignature.IDENTIFIER().getText());
            if (ctx.expr() == null && decl.getReturnType().getType() == Type.VOID) {
                this.getTopBody().getStatements().add(new ReturnStatement(null));
                return;
            } else if (decl.getReturnType().getType() == Type.VOID) {
                this.reporter.reportError(new ErrorReport(ctx.expr().start, "Attempt to return value from void function " + functionSignature.IDENTIFIER().getText()));
            }
            this.checkExprType(ctx.expr(), decl.getReturnType(), ctx.getStart(), this.reporter);
            this.getTopBody().getStatements().add(new ReturnStatement(this.getAstExpression(ctx.expr())));
        }
    }

    @Override
    public void enterIf_stmt(PocLang.If_stmtContext ctx) {
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            expectedIfExprType = new TypeContainer(Type.BOOLEAN, null);
            this.checkExprType(ctx.expr(), expectedIfExprType, ctx.expr().stop, this.reporter);


            IfStatement ifStmt = new IfStatement(getAstExpression(ctx.expr()), null, null, null);
            if (this.getTopBody().getParentStatement() instanceof IfStatement) {
                ((IfStatement) this.getTopBody().getParentStatement()).setElseIf(ifStmt);
            }
            this.currentStmt = ifStmt;
        }
    }

    @Override
    public void exitIf_stmt(PocLang.If_stmtContext ctx) {
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.currentStmt = null;
        }
    }

    @Override
    public void exitFunction_body(PocLang.Function_bodyContext ctx) {
        Stack<ParseTree> lines = new Stack<>();
        lines.add(ctx);
        boolean hasAtLeastOneReturn = false;

        while (!lines.empty()) {
            ParseTree pop = lines.pop();
            if (pop instanceof PocLang.Return_stmtContext) {
                hasAtLeastOneReturn = true;
            }
            for (int i = 0; i < pop.getChildCount(); i++) {
                lines.push(pop.getChild(i));
            }
        }
        ParserRuleContext parentCtx = ctx.getParent();
        while (!(parentCtx instanceof PocLang.FunctionContext)) parentCtx = parentCtx.getParent();
        PocLang.Function_sigContext functionSignature = (PocLang.Function_sigContext) ((PocLang.FunctionContext) parentCtx).children.get(0);
        FunctionDeclaration functionDeclaration = this.table.getFunctionByIdentifier(functionSignature.IDENTIFIER().getText());

        if (!hasAtLeastOneReturn && this.phase == VisitorPhase.COLLECTING_FUNCTIONS && functionDeclaration.getReturnType().getType() != Type.VOID) {
            String owningFunctionId = functionDeclaration.getIdentifier();
            this.reporter.reportError(new ErrorReport(ctx.getStop(), "No return statement in function " + owningFunctionId));
            return;
        }
    }

    @Override
    public void exitBody(PocLang.BodyContext ctx) {
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.scopeContainers.remove(this.scopeContainers.size() - 1);
        }
    }

    @Override
    public void enterBody(PocLang.BodyContext ctx) {
        // function sig visitor handles the very first body's scope container
        if (this.phase == VisitorPhase.CHECKING_TYPES && !(ctx.getParent() instanceof PocLang.Function_bodyContext)) {
            Body body = new Body(this.currentStmt);
            if (this.getTopBody().getParentStatement() instanceof IfStatement && ctx.getParent() instanceof PocLang.Optional_elseContext) {
                ((IfStatement)this.getTopBody().getParentStatement()).setElseBody(body);
            }

            this.scopeContainers.add(new ScopeContainer(body));
        }
    }

    private StackEntry lookupIdentifierInScopes(String identifier) {
        for (int i = this.scopeContainers.size() - 1; i >= 0; i--) {
            ScopeContainer scopeContainer = this.scopeContainers.get(i);
            if (scopeContainer.hasIdentifier(identifier)) {
                return scopeContainer.getByIdentifier(identifier);
            }
        }
        return null;
    }

    private Expression getAstExpression(PocLang.ExprContext exprCtx) {
        if (!exprCtx.expr().isEmpty()) {
            boolean isAdd = exprCtx.ADD() != null;
            boolean isMultiply = exprCtx.MULTIPLY() != null;
            boolean isDivide = exprCtx.DIVIDE() != null;
            boolean isSubtract = exprCtx.SUBTRACT() != null;
            BinaryOperationType type = (isAdd ? BinaryOperationType.INT_ADD : (isMultiply ? BinaryOperationType.INT_MULTIPLY : (
                    isDivide ? BinaryOperationType.INT_DIVIDE : BinaryOperationType.INT_SUBTRACT
            )));

            if (getExprType(exprCtx).getType() == Type.UNSIGNED_INTEGER) {
                return new BinaryOperationExpression(type, getAstExpression(exprCtx.expr(0)), getAstExpression(exprCtx.expr(1)));
            } else if (getExprType(exprCtx).getType() == Type.STRING && isAdd) {
                return new BinaryOperationExpression(BinaryOperationType.STRING_CONCAT, getAstExpression(exprCtx.expr(0)), getAstExpression(exprCtx.expr(1)));
            }
        } else if (exprCtx.function_call() != null) {
            PocLang.Function_callContext ctx = (PocLang.Function_callContext) exprCtx.function_call();
            return getAstExpressionForFunctionCall(ctx);
        } else if (exprCtx.value_ref() != null) {
            PocLang.Value_refContext value_refContext = exprCtx.value_ref();
            if (value_refContext.FALSE_LIT() != null || value_refContext.TRUE_LIT() != null) {
                return new BooleanLiteral(Boolean.parseBoolean(value_refContext.getText()), inferTypeFromValue(value_refContext));
            }
            if (value_refContext.INT() != null) {
                return new UIntLiteral(Long.parseLong(value_refContext.INT().getText()), inferTypeFromValue(value_refContext));
            }
            if (value_refContext.STRING_LITERAL() != null) {
                String text = value_refContext.STRING_LITERAL().getText();
                return new StringLiteral(value_refContext.STRING_LITERAL().getText().substring(1, text.length() - 1), inferTypeFromValue(value_refContext)); // ""
            } else if (value_refContext.identifier_ref() != null) {
                String text = value_refContext.identifier_ref().getText();
                StackEntry entry = lookupIdentifierInScopes(text);
                if (entry == null) {
                    return new ErrorNode();
                }
                return new VariableRefExpression(text, entry);
            }
        }
        throw new IllegalArgumentException("Can't handle this expression");
    }

    private FunctionCallExpression getAstExpressionForFunctionCall(PocLang.Function_callContext ctx) {
        List<Expression> arguments = new ArrayList<>();
        arguments.addAll(ctx.expr().stream().map(e -> getAstExpression(e)).collect(Collectors.toList()));
        return new FunctionCallExpression(this.table.getFunctionByIdentifier(ctx.IDENTIFIER().getText()), arguments);
    }
    
    private Body getTopBody() {
        return this.scopeContainers.get(this.scopeContainers.size()-1).getEnclosingBody();
    }
}