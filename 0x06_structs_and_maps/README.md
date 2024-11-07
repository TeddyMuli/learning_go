# Structs and Maps in Go
## Structs
  **Structs** are *user defined* types
  They allow for multiple data types
  e.g
  ```go
    type Person struct {
    name string;
    age int;
    weight float32;
    height float32;
  }
  ```

## Maps
  **Maps** allow for *key, value* pairs
  Use the map keyword
  Use **make** to create an empty map, if you dont use make it causes an runtime panic
  ```go
  make(map[string]string)
  ```
