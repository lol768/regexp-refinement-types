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