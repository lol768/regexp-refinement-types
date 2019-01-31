package eu.adamwilliams;

import eu.adamwilliams.reftypes.prototype.Application;
import eu.adamwilliams.reftypes.prototype.ErrorReport;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import eu.adamwilliams.reftypes.prototype.parser.PocLex;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.gradle.api.Plugin;
import org.gradle.api.Project;
import org.gradle.api.plugins.JavaPluginConvention;
import org.gradle.api.tasks.SourceSet;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class SimplePlugin implements Plugin<Project> {


    public void apply(Project project) {
        SourceSet mainSourceSet = project.getConvention().getPlugin(JavaPluginConvention.class).getSourceSets().getByName("main");
        File srcFolder = mainSourceSet.getAllSource().getSrcDirs().stream().map(d -> d.getParentFile()).findFirst().get();
        File srcRrt = new File(srcFolder, "rrt");

        project.task("poc", (a) -> {
            System.out.println("Performing type check");
            try (Stream<Path> paths = Files.walk(Paths.get(srcRrt.getAbsolutePath()))) {
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
                        System.err.println(report.getLineNo() + ":" + report.getLineNo() + " " + report.getMsg());
                    }
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        });
    }
}

