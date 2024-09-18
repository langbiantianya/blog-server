package routes

import (
	"blog-server/internal/entity"
	"blog-server/internal/entity/dto"
	"blog-server/internal/service"
	"blog-server/public/utils"
	"encoding/base64"
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
	var params dto.EssayDto
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.Error(err)
		return
	}
	res, err := essay.essayService.List(params)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (essay EssayRouter) Info(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}
	res, err := essay.essayService.Info(uint(uid))
	if err != nil {
		c.Error(err)
		return
	}

	if res.Post != "" {
		strb := base64.StdEncoding.EncodeToString([]byte(res.Post))
		res.Post = string(strb)
	}

	c.JSON(http.StatusOK, res)
}

func (essay EssayRouter) Add(c *gin.Context) {
	var params entity.Essay
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.Error(err)
		return
	}
	if params.Post != "" {
		strb, err := base64.StdEncoding.DecodeString(params.Post)
		if err != nil {
			c.Error(err)
			return
		}
		params.Post = string(strb)
	}
	err = essay.essayService.Add(params)
	if err != nil {
		c.Error(err)
		return
	}
	c.String(http.StatusOK, "OK")
}

func (essay EssayRouter) Update(c *gin.Context) {
	var params entity.Essay
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.Error(err)
		return
	}
	if params.Post != "" {
		strb, err := base64.StdEncoding.DecodeString(params.Post)
		if err != nil {
			c.Error(err)
			return
		}
		params.Post = string(strb)
	}

	err = essay.essayService.Update(params)

	if err != nil {
		c.Error(err)
		return
	}
	c.String(http.StatusOK, "OK")
}

func (essay EssayRouter) Hide(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}
	err = essay.essayService.Hide(uint(uid))
	if err != nil {
		c.Error(err)
		return
	}
	c.String(http.StatusOK, "OK")
}

func (essay EssayRouter) Delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}
	err = essay.essayService.Delete(uint(uid))
	if err != nil {
		c.Error(err)
		return
	}
	c.String(http.StatusOK, "OK")
}

func (essay EssayRouter) Publish(c *gin.Context) {
	ids := c.QueryArray("ids")
	uids := utils.Map(ids, func(index int, item string) (uint, error) {
		uid, err := strconv.ParseUint(item, 10, 32)
		return uint(uid), err
	})

	err := essay.essayService.Publish(uids)
	if err != nil {
		c.Error(err)
		return
	}
	c.String(http.StatusOK, "OK")
}
