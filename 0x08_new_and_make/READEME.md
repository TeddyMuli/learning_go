# New and Make in Go
The **new()** and **make()** keywords in Go are used to allocate memory but have different applications.

## new()
**new()** is used to allocate memory to  a new **zeroed** value data type like a *struct* and return its **pointer**.

e.g
```go
package main
import "fmt"

type Person struct {
  name string
  age int
}

func main() {
  person := new(Person)
  person.name = "Ted"
  person.age = 21

  fmt.Println(person)
}
```

## make()
**make()** is used to initialize data structures like slices and maps that need runtime initialization  and return a non-zeroed value of a specific type.
e.g
```go
package main

import "fmt"

func main() {
  slice1 := make([]int, 10, 15)

  for i := 0; i < 10; i++ {
    slice1[i] = i + 1
  }

  fmt.Println(slice1)
}
```
