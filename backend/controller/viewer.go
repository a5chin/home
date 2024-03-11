package controller

import (
	"github.com/gin-gonic/gin"
)

type ViewerController struct {
	ViewerUseCase
}

func NewViewerController(u ViewerUseCase) *ViewerController {
	return &ViewerController{ViewerUseCase: u}
}

type GetViewersResponse struct {
	Viewers uint `json:"viewers"`
}

// GetAViewers godoc
//
// @Summary    閲覧人数取得 API
// @Description
// @Tags       Viewer
// @Accept     json
// @Produce    json
// @Success    200     	{object}    GetViewersResponse
// @Failure    400     	{object}    entity.ErrorResponse
// @Failure    404     	{object}    entity.ErrorResponse
// @Router     /viewers [get]
func (c ViewerController) GetViewers(ctx *gin.Context) (interface{}, error) {
	viewers, err := c.ViewerUseCase.GetViewers(ctx)
	if err != nil {
		return nil, err
	}
	return GetViewersResponse{Viewers: viewers}, nil
}
