Prototype refinement types language
===================================

Example of syntax:

```
function LookupUserById(id: uint[> 1]): uint {
    return 1*1
}
```

Building
--------

Use the `build` Gradle task to build the project. As part of the Java source compilation, the grammar will be handled
by ANTLR and parser/lexer code generated.

You can run the application with the `run` task. Input is accepted from stdin.

Unit Tests
----------

Use the `test` task.

JavaSMT / z3
------------

Previously, this code made use of the JavaSMT abstraction layer (which is apathetic towards which underlying SMT solver is
used). For the sake of access to some string functionality, this dependency is being replaced with direct use of Z3.

Unfortunately, the Z3 Java bindings are not maintained on Maven Central. Thus, it is necessary to build Z3 manually.

See https://github.com/Z3Prover/z3/tree/cc312d2f688f4b6b4742fca12faedb1cb39e97b9#java for more information.