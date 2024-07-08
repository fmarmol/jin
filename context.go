package jin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/a-h/templ"
	"github.com/fmarmol/fp"
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

func (c Context) Get(key string) *Value {
	v, ok := c.Context.Get(key)
	return &Value{value: v, valid: ok}
}

func (v *Value) UUID() (uuid.UUID, error) {
	if v == nil || !v.valid {
		return uuid.Nil, ErrInvalidValue
	}
	switch _v := v.value.(type) {
	case string:
		return uuid.Parse(v.value.(string))
	case uuid.UUID:
		return _v, nil
	default:
		return uuid.Nil, fmt.Errorf("cannot convert type %T into uuid.UUID", reflect.TypeOf(v.value))
	}
}

func (v *Value) String() (string, error) {
	if v == nil || !v.valid {
		return "", ErrInvalidValue
	}
	return v.value.(string), nil
}

func (v *Value) Bool() (bool, error) {
	if v == nil || !v.valid {
		return false, ErrInvalidValue
	}
	return v.value.(bool), nil
}

func (c Context) GetCookie(name string) (*http.Cookie, error) {
	return c.Request.Cookie(name)
}

func (c Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Writer, cookie)
}

func (c Context) DeleteCookie(name string) error {
	cookie, err := c.GetCookie(name)
	if err != nil {
		return err
	}
	cookie.Expires = time.Now().AddDate(-1, 0, 0)
	c.SetCookie(cookie)
	return nil
}

func (c Context) ParseForm(dst any) error {
	err := c.Context.Request.ParseForm()
	if err != nil {
		return err
	}
	return fp.Parse(dst, c.Request.Form)
}

func (c Context) BadRequest(response any) (any, Error) {
	return nil, Error{http.StatusBadRequest, response}
}

func (c Context) InternalErr(response any) (any, Error) {
	return nil, Error{http.StatusInternalServerError, response}
}

func (c Context) Unauthorized(response any) (any, Error) {
	return nil, Error{http.StatusUnauthorized, response}
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
