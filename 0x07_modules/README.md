# Modules and Packages in Go
A package is a folder conatining a collection of related go source code

A module is a collection of packages with a go.mod file at its root that specifies version and dependencies

A package can also be a module if it has a go.mod file at its root

## Create a module
```bash
go mod init example.com/example
```
