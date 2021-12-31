package router

import (
	"net/http"

	"github.com/Fakorede/gin-app/controllers"
	"github.com/Fakorede/gin-app/dto"
	"github.com/gin-gonic/gin"
)

type VideoApis struct {
	videoController controllers.VideoController
}

func NewVideoAPIs(videoController controllers.VideoController) *VideoApis {
	return &VideoApis{
		videoController: videoController,
	}
}

// GetVideos godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Video
// @Failure 401 {object} dto.Response
// @Router /videos [get]
func (api *VideoApis) GetVideos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.videoController.FindAll())
}

// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos [post]
func (api *VideoApis) CreateVideo(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

// UpdateVideo godoc
// @Security bearerAuth
// @Summary Update videos
// @Description Update a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Param video body entity.Video true "Update video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [put]
func (api *VideoApis) UpdateVideo(ctx *gin.Context) {
	err := api.videoController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

// DeleteVideo godoc
// @Security bearerAuth
// @Summary Remove videos
// @Description Delete a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [delete]
func (api *VideoApis) DeleteVideo(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

