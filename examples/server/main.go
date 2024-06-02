package main

import (
	"fmt"
	"log"

	"github.com/fmarmol/jin"
)

func Hello(c jin.Context) (any, error) {
	v := c.GetString("myKey")
	if v != "" {
		fmt.Println("HERE the value", v)
	}
	return c.BadRequest("request is not valid")
}
func World(c jin.Context) (any, error) {
	return "test", nil
}

func aMiddleware(c jin.Context) (any, error) {
	log.Println("a middleware")
	c.Set("myKey", "value")
	return nil, nil
}

func main() {
	e := jin.New()
	e.GET("/hello", Hello)
	e.GET("/world", World)
	g := e.Group("/books", aMiddleware)
	g.GET("/hello", Hello)
	log.Fatal(e.Run(":8082"))
}
