package eu.adamwilliams.reftypes.prototype;

import com.microsoft.z3.BoolExpr;
import com.microsoft.z3.Context;
import com.microsoft.z3.SeqExpr;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;

public interface IRegexZ3Adapter {
    BoolExpr getRefinementType(SeqExpr y, PocLang.ReContext regExp, Context z3Ctx);
}
