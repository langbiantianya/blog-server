package routes

import (
	"blog-server/internal/entity"
	"blog-server/internal/entity/dto"
	"blog-server/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IEssayRouter interface {
	Info(*gin.Context)
	List(*gin.Context)
	Add(*gin.Context)
	Update(*gin.Context)
	Hide(*gin.Context)
	Delete(*gin.Context)
	Publish(*gin.Context)
}

type EssayRouter struct {
	essayService service.IEssayService
}

func NewEssayRouter(essayService service.IEssayService, tagService service.ITagService) IEssayRouter {
	return &EssayRouter{
		essayService: essayService,
	}
}

func (essay EssayRouter) List(c *gin.Context) {
	var params dto.EssayDTO
	err := c.ShouldBindQuery(&params)
	if err != nil {
		panic(err)

	}
	res, err := essay.essayService.List(params)
	if err != nil {
		panic(err)

	}

	c.JSON(http.StatusOK, res)
}

func (essay EssayRouter) Info(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err)

	}
	res, err := essay.essayService.Info(uint(uid))
	if err != nil {
		panic(err)

	}

	c.JSON(http.StatusOK, res)
}

func (essay EssayRouter) Add(c *gin.Context) {
	var params entity.Essay
	err := c.ShouldBindJSON(&params)
	if err != nil {
		panic(err)

	}
	// if params.Post != "" {
	// 	strb, err := base64.StdEncoding.DecodeString(params.Post)
	// 	if err != nil {
	// 		panic(err)
	//
	// 	}
	// 	params.Post = string(strb)
	// }
	err = essay.essayService.Add(params)
	if err != nil {
		panic(err)

	}
	c.String(http.StatusOK, "OK")
}

func (essay EssayRouter) Update(c *gin.Context) {
	var params entity.Essay
	err := c.ShouldBindJSON(&params)
	if err != nil {
		panic(err)

	}
	// if params.Post != "" {
	// 	strb, err := base64.StdEncoding.DecodeString(params.Post)
	// 	if err != nil {
	// 		panic(err)
	//
	// 	}
	// 	params.Post = string(strb)
	// }

	err = essay.essayService.Update(params)

	if err != nil {
		panic(err)

	}
	c.String(http.StatusOK, "OK")
}

func (essay EssayRouter) Hide(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err)

	}
	err = essay.essayService.Hide(uint(uid))
	if err != nil {
		panic(err)

	}
	c.String(http.StatusOK, "OK")
}

func (essay EssayRouter) Delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err)

	}
	err = essay.essayService.Delete(uint(uid))
	if err != nil {
		panic(err)

	}
	c.String(http.StatusOK, "OK")
}

func (essay EssayRouter) Publish(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err)

	}
	err = essay.essayService.Publish(uint(uid))
	if err != nil {
		panic(err)

	}
	c.String(http.StatusOK, "OK")
}
