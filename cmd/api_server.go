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
	v1 := r.Group("/api/v1")
	essay := v1.Group("/essay")
	essay.GET("/:id", api.essay.Info)
	essay.GET("/list", api.essay.List)
	essay.POST("/", api.essay.Add)
	essay.PUT("/", api.essay.Update)
	essay.DELETE("/:id", api.essay.Delete)
	essay.PUT("/hide/:id", api.essay.Hide)
	essay.PUT("/publish", api.essay.Publish)
}
