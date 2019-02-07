package eu.adamwilliams;

import org.gradle.api.Plugin;
import org.gradle.api.Project;
import org.gradle.api.plugins.JavaPluginConvention;
import org.gradle.api.tasks.SourceSet;

import java.io.File;


public class TypeCheckPlugin implements Plugin<Project> {


    public void apply(Project project) {
        SourceSet mainSourceSet = project.getConvention().getPlugin(JavaPluginConvention.class).
                getSourceSets().getByName("main");
        File srcFolder = mainSourceSet.getAllSource().getSrcDirs().stream().map(d -> d.getParentFile()).findFirst().get();
        File srcRrt = new File(srcFolder, "rrt");

        project.getTasks().create("poc", TypeCheckTask.class, (a) -> {
            a.setSrcRt(srcRrt);
            a.dependsOn("compileJava");
        });
    }
}

