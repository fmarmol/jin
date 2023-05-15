package ginjson

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (c *Context) BadRequest(response any) (any, Error) {
	return nil, Error{http.StatusBadRequest, response}
}
func (c *Context) InternalErr(response any) (any, Error) {
	return nil, Error{http.StatusInternalServerError, response}
}
func NotFound(response any) (any, Error) {
	return nil, Error{http.StatusNotFound, response}
}
