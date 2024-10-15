package middleware

import (
	"blog-server/internal/entity/vo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	err := c.Err()
	if err != nil && !c.IsAborted() {
		c.AbortWithStatusJSON(http.StatusBadRequest, vo.Fail(c.Errors.Last()))
	} else if c.Errors != nil && len(c.Errors) > 0 && !c.IsAborted() {
		c.AbortWithStatusJSON(http.StatusInternalServerError, vo.Fail(err))
	} else if c.Errors != nil && len(c.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, vo.Fail(fmt.Errorf("%v", err)))
	}
}
func PanicRecovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, vo.Fail(fmt.Errorf("%v", err)))
		}
	}()
	c.Next()
}
