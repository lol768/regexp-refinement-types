//Program to be parsed
String program = "function Main()\n{\n......";
// Lexer for the language, pointed at a stream constructed
// from the string above
PocLex lexer = new PocLex(CharStreams.fromString(program));
CommonTokenStream tokens = new CommonTokenStream(lexer);
// Parser, depends on tokens lexed by lexer
PocLang parser = new PocLang(tokens);
// Do the parsing, starting with the "program" non-terminal
ProgramContext parsedProgram = parser.program();