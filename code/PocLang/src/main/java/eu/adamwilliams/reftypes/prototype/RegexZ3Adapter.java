package eu.adamwilliams.reftypes.prototype;

import com.microsoft.z3.BoolExpr;
import com.microsoft.z3.Context;
import com.microsoft.z3.ReExpr;
import com.microsoft.z3.SeqExpr;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import org.antlr.v4.runtime.ParserRuleContext;

public class RegexZ3Adapter implements IRegexZ3Adapter {
    @Override
    public BoolExpr getRefinementType(SeqExpr y, PocLang.ReContext regExp, Context z3Ctx) {
        return z3Ctx.mkInRe(y, auxiliaryExpression(regExp, z3Ctx));
    }

    private ReExpr auxiliaryExpression(ParserRuleContext parserObj, Context z3Ctx) {
        if (parserObj instanceof PocLang.ReContext) {
            PocLang.Simple_reContext simpleRe = ((PocLang.ReContext) parserObj).simple_re();
            PocLang.Union_primeContext unionPrime = ((PocLang.ReContext) parserObj).union_prime();
            if (unionPrime.re() == null) {
                return auxiliaryExpression(simpleRe, z3Ctx);
            }
            return z3Ctx.mkUnion(auxiliaryExpression(simpleRe, z3Ctx), auxiliaryExpression(unionPrime.re(), z3Ctx));
        } else if (parserObj instanceof PocLang.Simple_reContext) {
            PocLang.Simple_reContext simpleRe = (PocLang.Simple_reContext) parserObj;
            if (simpleRe.concat_prime().simple_re() == null) {
                return auxiliaryExpression(simpleRe.basic_re(), z3Ctx);
            }
            return z3Ctx.mkConcat(auxiliaryExpression(simpleRe.basic_re(), z3Ctx), auxiliaryExpression(simpleRe.concat_prime().simple_re(), z3Ctx));
        } else if (parserObj instanceof PocLang.Basic_reContext) {
            PocLang.Basic_reContext basic = (PocLang.Basic_reContext) parserObj;
            if (basic.plus() != null) {
                return z3Ctx.mkPlus(auxiliaryExpression(basic.plus().elementary_re(), z3Ctx));
            }

            if (basic.kleene_star() != null) {
                return z3Ctx.mkStar(auxiliaryExpression(basic.kleene_star().elementary_re(), z3Ctx));
            }

            if (basic.elementary_re() != null) {
                return auxiliaryExpression(basic.elementary_re(), z3Ctx);
            }
        } else if (parserObj instanceof PocLang.Elementary_reContext) {
            PocLang.Elementary_reContext elementary = (PocLang.Elementary_reContext) parserObj;
            if (elementary.character() != null) {
                return z3Ctx.mkToRe(z3Ctx.mkString(elementary.character().getText()));
            }

            if (elementary.DOT() != null) {
                return z3Ctx.mkRange(z3Ctx.mkString("\u0000"), z3Ctx.mkString("\uffff"));
            }

            if (elementary.range() != null) {
                PocLang.RangeContext items = elementary.range();
                if (items.positive_range() != null) {
                    PocLang.Positive_rangeContext positive_rangeContext = items.positive_range();
                    return handleRegexRange(positive_rangeContext.range_items(), z3Ctx);
                } else {
                    PocLang.Negative_rangeContext negative_rangeContext = items.negative_range();
                    return z3Ctx.mkComplement(handleRegexRange(negative_rangeContext.range_items(), z3Ctx));
                }
            }

            if (elementary.group() != null) {
                return auxiliaryExpression(elementary.group().re(), z3Ctx);
            }
        }
        throw new IllegalArgumentException("Unacceptable parser object");
    }

    private ReExpr handleRegexRange(PocLang.Range_itemsContext staged, Context z3Ctx) {
        ReExpr expr = null;
        while (staged != null) {
            PocLang.Range_itemContext item = staged.range_item();
            ReExpr stagedExpr;
            if (item.MINUS() != null) {
                stagedExpr = z3Ctx.mkRange(z3Ctx.mkString(item.lax_character(0).getText()), z3Ctx.mkString(item.lax_character(1).getText()));
            } else {
                stagedExpr = z3Ctx.mkRange(z3Ctx.mkString(item.lax_character(0).getText()), z3Ctx.mkString(item.lax_character(0).getText()));
            }
            expr = (expr != null) ? z3Ctx.mkUnion(expr, stagedExpr) : stagedExpr;
            staged = staged.range_items();
        }

        return expr;
    }
}


