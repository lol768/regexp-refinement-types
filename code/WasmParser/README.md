```
antlr4 -Dlanguage=Go -visitor -o /home/adam/Documents/Go/src/WasmParser/parser/ PocLex.g4
antlr4 -Dlanguage=Go -visitor -o /home/adam/Documents/Go/src/WasmParser/parser/ PocLang.g4
```

```
GOOS=js GOARCH=wasm go build -o Test.wasm Test.go
```