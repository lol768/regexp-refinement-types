package eu.adamwilliams.reftypes.prototype;

import org.antlr.v4.runtime.Token;

import java.util.List;

public interface ErrorReporter {
    void reportError(ErrorReport report);
    List<ErrorReport> getReports();

}

