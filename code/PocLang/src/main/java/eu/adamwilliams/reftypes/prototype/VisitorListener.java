package eu.adamwilliams.reftypes.prototype;

import eu.adamwilliams.reftypes.prototype.domain.FunctionDeclaration;
import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.TypeContainer;
import eu.adamwilliams.reftypes.prototype.domain.VisitorPhase;
import eu.adamwilliams.reftypes.prototype.parser.PocLangBaseListener;
import eu.adamwilliams.reftypes.prototype.parser.PocLangParser;
import org.antlr.v4.runtime.Token;

import java.util.HashMap;
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
    private VisitorPhase phase = VisitorPhase.COLLECTING_FUNCTIONS;
    private Map<String, TypeContainer> typesCurrentlyInScope = new HashMap<>();

    public VisitorListener(FunctionTable table, ErrorReporter reporter) {
        this.table = table;
        this.reporter = reporter;
    }

    public void setPhase(VisitorPhase phase) {
        this.phase = phase;
    }


    @Override
    public void enterFunction_sig(PocLangParser.Function_sigContext ctx) {
        super.enterFunction_sig(ctx);
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
        } else if (this.phase == VisitorPhase.CHECKING_TYPES) {
            Map<String, TypeContainer> collect = ctx.argument_decl().stream().collect(Collectors.toMap(ad -> ad.IDENTIFIER().getText(), ad -> mapTypeFromParsed(ad.type())));
            this.table.addFunction(new FunctionDeclaration(idText, mapTypeFromParsed(ctx.type()), collect, symbol.getLine(), symbol.getCharPositionInLine()));

            this.typesCurrentlyInScope.putAll(collect);
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

        return new TypeContainer(type, ctx.int_constraint() != null ? ctx.int_constraint().getText() : null);

    }

    @Override
    public void enterVar_decl(PocLangParser.Var_declContext ctx) {
        super.enterVar_decl(ctx);
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.typesCurrentlyInScope.put(ctx.IDENTIFIER().getText(), mapTypeFromParsed(ctx.type()));
        }
    }

    @Override
    public void enterFunction_call(PocLangParser.Function_callContext ctx) {
        super.enterFunction_call(ctx);
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            String idText = ctx.IDENTIFIER().getText();
            Token symbol = ctx.IDENTIFIER().getSymbol();

            if (!this.table.hasFunction(idText)) {
                String msg = String.format("Call to undefined function with identifier %s", idText);
                this.reporter.reportError(new ErrorReport(symbol, msg));
                return;
            }

            // check types of arguments
        }
    }

    @Override
    public void exitBody(PocLangParser.BodyContext ctx) {
        List<PocLangParser.Body_lineContext> lines = ctx.body_line();
        if (!lines.stream().flatMap(bl -> bl.children.stream()).anyMatch(line -> line instanceof PocLangParser.Return_stmtContext) && this.phase == VisitorPhase.COLLECTING_FUNCTIONS) {
            String owningFunctionId = ((PocLangParser.Function_sigContext) ctx.getParent().children.get(0)).IDENTIFIER().getText();
            this.reporter.reportError(new ErrorReport(lines.get(lines.size() - 1).getStart(), "No return statement in function " + owningFunctionId));
            return;
        }
        if (this.phase == VisitorPhase.CHECKING_TYPES) {
            this.typesCurrentlyInScope.clear();
        }
    }
}