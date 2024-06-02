# Documentation
This `jin` package is a simple wrapper of the gin framework mainly focus on json response type and remove boiler plates

## Middleware
just return nil for both response and error

```go
func myMiddleware(c Context) (any, error) {
  return nil,nil
}
````
