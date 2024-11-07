# Defer in Go
Defer delays the execution of a function till the surrounding function is executed.

If you have multiple defer statements the first defer statement will be the last to be executed.
e.g
```go
defer fmt.Println("world!")
defer fmt.Println("Ted")
fmt.Println("Hello")
```

The output will be as follows
```bash
Hello
Ted
world!
```

The first defer statement delays till the next function is executed, then the 2nd defer statement delays till the 3rd function id executed.
When the third function is executed, the 2nd function which was dependenet on the execution of the 3rd function is executed then the 1st function which was dependent on the 2nd function is executed
