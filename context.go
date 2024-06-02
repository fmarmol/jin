package jin

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (c Context) BadRequest(response any) (any, Error) {
	return nil, Error{http.StatusBadRequest, response}
}

func (c Context) InternalErr(response any) (any, Error) {
	return nil, Error{http.StatusInternalServerError, response}
}

func (c Context) NotFound(response any) (any, Error) {
	return nil, Error{http.StatusNotFound, response}
}

func (c Context) HtmxRefresh() (any, error) {
	c.Header("HX-Refresh", "true")
	return nil, nil
}

func (c Context) HtmxRedirect(location string) (any, error) {
	c.Header("HX-Redirect", location)
	return nil, nil
}

func (c Context) TemplWithCtx(ctx context.Context, component templ.Component) (any, error) {
	err := component.Render(ctx, c.Writer)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c Context) Templ(component templ.Component) (any, error) {
	return c.TemplWithCtx(c, component)
}
