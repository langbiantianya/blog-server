package cmd

import (
	"blog-server/internal/middleware"
	"blog-server/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiRoutes struct {
	essay routes.IEssayRouter
	tag   routes.ITagRouter
}

func StartApiServer(port int, apiRoutes *ApiRoutes) {
	r := gin.Default()
	r.Use(middleware.PanicRecovery)
	apiRoutes.apiV1(r)
	r.Run(fmt.Sprintf(":%d", port))
}

func NewApiRoutes(essay routes.IEssayRouter, tag routes.ITagRouter) *ApiRoutes {
	return &ApiRoutes{
		essay: essay,
		tag:   tag,
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
	essay.PUT("/publish/:id", api.essay.Publish)
	tag := v1.Group("/tag")
	tag.GET("/list", api.tag.List)
}
