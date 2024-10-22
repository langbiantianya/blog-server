package middleware

import (
	"blog-server/internal/entity/vo"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	err := c.Err()
	if err != nil && !c.IsAborted() {
		log.Fatal(err, string(debug.Stack()))
		c.AbortWithStatusJSON(http.StatusBadRequest, vo.Fail(c.Errors.Last()))
	} else if c.Errors != nil && len(c.Errors) > 0 && !c.IsAborted() {
		log.Fatal(err, string(debug.Stack()))
		c.AbortWithStatusJSON(http.StatusInternalServerError, vo.Fail(err))
	} else if c.Errors != nil && len(c.Errors) > 0 {
		log.Fatal(err, string(debug.Stack()))
		c.AbortWithStatusJSON(http.StatusInternalServerError, vo.Fail(fmt.Errorf("%v", err)))
	}
}
func PanicRecovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, string(debug.Stack()))
			c.AbortWithStatusJSON(http.StatusInternalServerError, vo.Fail(fmt.Errorf("%v", err)))
		}
	}()
	c.Next()
}
