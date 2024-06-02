package main

import (
	"fmt"
	"log"

	gin "github.com/fmarmol/ginjson"
)

func Hello(c gin.Context) (any, error) {
	v := c.GetString("myKey")
	if v != "" {
		fmt.Println("HERE the value", v)
	}
	return c.BadRequest("request is not valid")
}
func World(c gin.Context) (any, error) {
	return "test", nil
}

func aMiddleware(c gin.Context) (any, error) {
	log.Println("a middleware")
	c.Set("myKey", "value")
	return nil, nil
}

func main() {
	e := gin.New()
	e.GET("/hello", Hello)
	e.GET("/world", World)
	g := e.Group("/books", aMiddleware)
	g.GET("/hello", Hello)
	log.Fatal(e.Run(":8080"))
}
