# Go Structs to TS Types Transpiler

This repo uses Go's built-in `go/ast`, `go/parser`, and `go/token` to convert files containing Go structs to their TypeScript type equivalent.

To try it out, run the following command:
```
go run cmd/*.go
```

A TODO is to write the transpiled structs into a new file instead of logging to stdout.