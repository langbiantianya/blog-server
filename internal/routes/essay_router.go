package routes

import (
	"blog-server/internal/service"

	"github.com/gin-gonic/gin"
)

type IEssayRouter interface {
	Info(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Hide(c *gin.Context)
	Delete(c *gin.Context)
}

type EssayRouter struct {
	essayService service.IEssayService
}

func NewEssayRouter(essayService service.IEssayService) IEssayRouter {
	return EssayRouter{
		essayService: essayService,
	}
}

func (essay EssayRouter) Info(c *gin.Context) {
	panic("TODO")
}

func (essay EssayRouter) Add(c *gin.Context) {
	panic("TODO")

}

func (essay EssayRouter) Update(c *gin.Context) {
	panic("TODO")

}

func (essay EssayRouter) Hide(c *gin.Context) {
	panic("TODO")

}

func (essay EssayRouter) Delete(c *gin.Context) {
	panic("TODO")

}
