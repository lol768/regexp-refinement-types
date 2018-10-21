package eu.adamwilliams.reftypes.prototype;

import eu.adamwilliams.reftypes.prototype.domain.FunctionDeclaration;
import eu.adamwilliams.reftypes.prototype.domain.Type;
import eu.adamwilliams.reftypes.prototype.domain.VisitorPhase;
import eu.adamwilliams.reftypes.prototype.parser.PocLangBaseListener;
import eu.adamwilliams.reftypes.prototype.parser.PocLangParser;
import org.antlr.v4.runtime.Parser;
import org.antlr.v4.runtime.RecognitionException;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.tree.ErrorNodeImpl;

/**
 * Visitor for items in the parse tree. We'll do things
 * like register functions in the function table here.
 */
public class VisitorListener extends PocLangBaseListener {
    private FunctionTable table;
    private ErrorReporter reporter;
    private VisitorPhase phase = VisitorPhase.COLLECTING_FUNCTIONS;

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
        if (this.phase == VisitorPhase.COLLECTING_FUNCTIONS) {
            String idText = ctx.IDENTIFIER().getText();
            Token symbol = ctx.IDENTIFIER().getSymbol();
            if (this.table.hasFunction(idText)) {
                FunctionDeclaration function = this.table.getFunctionByIdentifier(idText);
                String prevLocation = String.format("L%d:%d", function.getLineNo(), function.getColNo());
                String msg = String.format("Redeclaration of function %s, previously declared at %s", ctx.IDENTIFIER(), prevLocation);
                this.reporter.reportError(new ErrorReport(symbol, msg));
                return;
            }

            this.table.addFunction(new FunctionDeclaration(idText, Type.UNSIGNED_INTEGER, symbol.getLine(), symbol.getCharPositionInLine()));
        }
    }

    @Override
    public void enterFunction_call(PocLangParser.Function_callContext ctx) {
        super.enterFunction_call(ctx);
        if (this.phase == VisitorPhase.CHECKING_FUNCTION_CALLS) {
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

    // TODO: enter function bodies, create stack for locals and store their types
}