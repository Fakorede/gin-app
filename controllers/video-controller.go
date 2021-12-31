package controllers

import (
	"net/http"
	"strconv"

	"github.com/Fakorede/gin-app/entity"
	"github.com/Fakorede/gin-app/services"
	"github.com/Fakorede/gin-app/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type videoController struct {
	service services.VideoService
}

var validate *validator.Validate

func NewVideoController(service services.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("custom", validators.ValidateCustomTitle)

	return &videoController{
		service: service,
	}
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *videoController)	Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}

func (c *videoController) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Update(video)
	return nil
}

func (c *videoController) Delete(ctx *gin.Context) error {
	var video entity.Video

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id

	c.service.Delete(video)
	return nil
}

func (c *videoController) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title": "All Videos",
		"videos": videos,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}