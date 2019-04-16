public class TypeCheckPlugin implements Plugin<Project> {
    public void apply(Project project) {
        SourceSet mainSources = project.getConvention()
                .getPlugin(JavaPluginConvention.class)
                .getSourceSets().getByName("main"); //get main source set

        // within main, look for an `rrt` directory
        File srcDir = mainSources.getAllSource().getSrcDirs().stream().map(
                d -> d.getParentFile()
        ).findFirst().get();
        File srcRrt = new File(srcDir, "rrt");

        // set up the task so it happens *after* Java has compiled, for the FFI
        project.getTasks().create("rrt", TypeCheckTask.class, (a) -> {
            a.setSrcRt(srcRrt);
            a.dependsOn("compileJava");
        });
    }
}