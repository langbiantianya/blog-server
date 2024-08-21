package cmd

import (
	"blog-server/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiRoutes struct {
	essay routes.IEssayRouter
}

func StartApiServer(port int, apiRoutes *ApiRoutes) {
	r := gin.Default()
	apiRoutes.apiV1(r)
	r.Run(fmt.Sprintf(":%d", port))
}

func NewApiRoutes(essay routes.IEssayRouter) *ApiRoutes {
	return &ApiRoutes{
		essay: essay,
	}
}

func (api ApiRoutes) apiV1(r *gin.Engine) {
	var v1 = r.Group("/api/v1")
	v1.GET("/essay/:id", api.essay.Info)
}
