package controllers

import (
	"github.com/Fakorede/gin-app/entity"
	"github.com/Fakorede/gin-app/services"
	"github.com/Fakorede/gin-app/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type videoController struct {
	service services.VideoService
}

var validate *validator.Validate

func New(service services.VideoService) VideoController {
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