package eu.adamwilliams.reftypes.prototype;

import java.util.ArrayList;
import java.util.List;

public class ErrorReporterImpl implements ErrorReporter {

    private List<ErrorReport> reports = new ArrayList<>();

    @Override
    public void reportError(ErrorReport report) {
        this.reports.add(report);
    }

    @Override
    public List<ErrorReport> getReports() {
        return this.reports;
    }
}

