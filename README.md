# Documentation
This `ginjson` package is a simple wrapper gin framework to use on api with json response type and remove boiler plates

## Middleware
just return nil for both response and error

```go
func myMiddleware(c *Context) (any, error) {
  return nil,nil
}
````
# ginjson
