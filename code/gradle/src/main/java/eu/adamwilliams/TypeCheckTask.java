package eu.adamwilliams;

import eu.adamwilliams.reftypes.prototype.Application;
import eu.adamwilliams.reftypes.prototype.ErrorReport;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import eu.adamwilliams.reftypes.prototype.parser.PocLex;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.gradle.api.DefaultTask;
import org.gradle.api.GradleException;
import org.gradle.api.tasks.TaskAction;
import org.gradle.internal.logging.text.StyledTextOutput;
import org.gradle.internal.logging.text.StyledTextOutputFactory;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class TypeCheckTask extends DefaultTask {
    private File srcRt;

    public File getSrcRt() {
        return srcRt;
    }

    public void setSrcRt(File srcRt) {
        this.srcRt = srcRt;
    }

    @TaskAction
    void doTypeCheck() {
        StyledTextOutput out = this.getServices().get(StyledTextOutputFactory.class).create("typeCheck");
        System.out.println(System.getProperty("java.library.path"));
        boolean fail = false;
        try (Stream<Path> paths = Files.walk(Paths.get(this.srcRt.getAbsolutePath()))) {
            List<Path> list = paths
                    .filter(Files::isRegularFile)
                    .collect(Collectors.toList());
            for (Path path : list) {
                PocLex lexer = new PocLex(CharStreams.fromFileName(path.toAbsolutePath().toString()));
                CommonTokenStream tokens = new CommonTokenStream(lexer);
                PocLang parser = new PocLang(tokens);
                PocLang.ProgramContext program = parser.program();
                Application app = new Application();
                List<ErrorReport> reports = app.doTypeChecks(program).getReports();

                for (ErrorReport report : reports) {
                    fail = true;
                    out.withStyle(StyledTextOutput.Style.Error).println(path.toAbsolutePath().toString().replace(this.srcRt.getAbsolutePath()+"/", "") + ": L" + report.getLineNo() + ":" + report.getLineNo() + " " + report.getMsg());
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        if (fail) {
            throw new GradleException("The type check failed!");
        }
    }
}
