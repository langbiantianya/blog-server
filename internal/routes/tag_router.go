package routes

import (
	"blog-server/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ITagRouter interface {
	List(*gin.Context)
}

type TagRouter struct {
	tagService service.ITagService
}

func NewTagRouter(tagService service.ITagService) ITagRouter {
	return TagRouter{
		tagService,
	}
}

func (tag TagRouter) List(c *gin.Context) {
	res, err := tag.tagService.List()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}
