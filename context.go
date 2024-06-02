package jin

import (
	"context"
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Context struct {
	*gin.Context
}

type Value struct {
	value any
	valid bool
}

var ErrInvalidValue = errors.New("invalid value")

func (v *Value) UUID() (uuid.UUID, error) {
	if !v.valid {
		return uuid.Nil, ErrInvalidValue
	}
	return uuid.Parse(v.value.(string))
}

func (v *Value) String() (string, error) {
	if !v.valid {
		return "", ErrInvalidValue
	}
	return v.value.(string), nil
}

func (c Context) Get(key string) *Value {
	v, ok := c.Context.Get(key)
	return &Value{value: v, valid: ok}
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

func (c Context) Redirect(location string) (any, error) {
	c.Context.Redirect(http.StatusSeeOther, location)
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
