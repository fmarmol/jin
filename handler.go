package jin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler func(c Context) (any, error)

func (h Handler) ToGin() func(c *gin.Context) {
	return func(c *gin.Context) {
		resp, err := h(Context{c})
		if err == nil {
			if resp != nil {
				c.JSON(http.StatusOK, resp)
				return
			}
		} else {
			c.Error(err)
			switch e := err.(type) {
			case Error:
				c.JSON(e.Code, e.Response)
				c.Abort()
				return
			default: // fallback for generic error
				c.JSON(http.StatusInternalServerError, e.Error())
				c.Abort()
				return
			}
		}
	}
}
