package eu.adamwilliams.reftypes.prototype;

import org.antlr.v4.runtime.Token;

import java.util.List;

public interface ErrorReporter {
    void reportError(ErrorReport report);
    List<ErrorReport> getReports();

}

class ErrorReport {
    private Token token;
    private String msg;

    public ErrorReport(Token token, String msg) {
        this.token = token;
        this.msg = msg;
    }

    public Token getToken() {
        return token;
    }

    public String getMsg() {
        return msg;
    }
}