package jin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func (e *RouterGroup) add(method string, path string, handlers ...Handler) {

	hs := make([]gin.HandlerFunc, 0, len(handlers))
	for _, h := range handlers {
		hs = append(hs, h.ToGin())
	}
	switch method {
	case "HEAD":
		e.RouterGroup.HEAD(path, hs...)
	case "GET":
		e.RouterGroup.GET(path, hs...)
	case "PUT":
		e.RouterGroup.PUT(path, hs...)
	case "DELETE":
		e.RouterGroup.DELETE(path, hs...)
	case "POST":
		e.RouterGroup.POST(path, hs...)
	case "PATCH":
		e.RouterGroup.PATCH(path, hs...)
	case "OPTIONS":
		e.RouterGroup.OPTIONS(path, hs...)
	default:
		panic(fmt.Errorf("unknown method %v", method))
	}
}

func (e *RouterGroup) Head(path string, handlers ...Handler)    { e.add("HEAD", path, handlers...) }
func (e *RouterGroup) GET(path string, handlers ...Handler)     { e.add("GET", path, handlers...) }
func (e *RouterGroup) PUT(path string, handlers ...Handler)     { e.add("PUT", path, handlers...) }
func (e *RouterGroup) DELETE(path string, handlers ...Handler)  { e.add("DELETE", path, handlers...) }
func (e *RouterGroup) POST(path string, handlers ...Handler)    { e.add("POST", path, handlers...) }
func (e *RouterGroup) PATCH(path string, handlers ...Handler)   { e.add("PATCH", path, handlers...) }
func (e *RouterGroup) OPTIONS(path string, handlers ...Handler) { e.add("OPTIONS", path, handlers...) }
