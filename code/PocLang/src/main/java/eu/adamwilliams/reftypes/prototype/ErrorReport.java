package eu.adamwilliams.reftypes.prototype;

import org.antlr.v4.runtime.Token;

public class ErrorReport {
    private transient Token token;
    private String msg;
    private int lineNo;
    private int colNo;

    public ErrorReport(Token token, String msg) {
        this.token = token;
        this.msg = msg;
        this.lineNo = token.getLine();
        this.colNo = token.getCharPositionInLine();
    }

    public Token getToken() {
        return token;
    }

    public String getMsg() {
        return msg;
    }

    public int getLineNo() {
        return lineNo;
    }

    public int getColNo() {
        return colNo;
    }
}
