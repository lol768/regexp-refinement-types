package eu.adamwilliams.reftypes.prototype.domain;

import eu.adamwilliams.reftypes.prototype.ErrorReport;
import eu.adamwilliams.reftypes.prototype.ErrorReporter;
import eu.adamwilliams.reftypes.prototype.FunctionTable;

import java.util.List;

public class TypeCheckResults {
    private ErrorReporter errorReporter;
    private FunctionTable functionTable;

    public TypeCheckResults(ErrorReporter errorReporter, FunctionTable functionTable) {
        this.errorReporter = errorReporter;
        this.functionTable = functionTable;
    }

    public FunctionTable getFunctionTable() {
        return functionTable;
    }

    public List<ErrorReport> getReports() {
        return this.errorReporter.getReports();
    }
}
