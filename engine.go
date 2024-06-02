package jin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
}

func SetRealeaseMode() {
	gin.SetMode(gin.ReleaseMode)
}

func SetDebugMode() {
	gin.SetMode(gin.DebugMode)
}

func New() Engine {
	return Engine{Engine: gin.New()}
}

func (e Engine) add(method string, path string, handlers ...Handler) {

	hs := make([]gin.HandlerFunc, 0, len(handlers))
	for _, h := range handlers {
		hs = append(hs, h.ToGin())
	}
	switch method {
	case "HEAD":
		e.Engine.HEAD(path, hs...)
	case "GET":
		e.Engine.GET(path, hs...)
	case "PUT":
		e.Engine.PUT(path, hs...)
	case "DELETE":
		e.Engine.DELETE(path, hs...)
	case "POST":
		e.Engine.POST(path, hs...)
	case "PATCH":
		e.Engine.PATCH(path, hs...)
	case "OPTIONS":
		e.Engine.OPTIONS(path, hs...)
	default:
		panic(fmt.Errorf("unknown method %v", method))
	}
}

func (e Engine) HEAD(path string, handlers ...Handler)    { e.add("HEAD", path, handlers...) }
func (e Engine) GET(path string, handlers ...Handler)     { e.add("GET", path, handlers...) }
func (e Engine) PUT(path string, handlers ...Handler)     { e.add("PUT", path, handlers...) }
func (e Engine) DELETE(path string, handlers ...Handler)  { e.add("DELETE", path, handlers...) }
func (e Engine) POST(path string, handlers ...Handler)    { e.add("POST", path, handlers...) }
func (e Engine) PATCH(path string, handlers ...Handler)   { e.add("PATCH", path, handlers...) }
func (e Engine) OPTIONS(path string, handlers ...Handler) { e.add("OPTIONS", path, handlers...) }

func (e Engine) Group(path string, handlers ...Handler) *RouterGroup {
	hs := make([]gin.HandlerFunc, 0, len(handlers))
	for _, h := range handlers {
		hs = append(hs, h.ToGin())
	}
	return &RouterGroup{e.Engine.Group(path, hs...)}
}
