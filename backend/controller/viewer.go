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

type GetTotalViewersResponse struct {
	Viewers uint `json:"viewers"`
}

// GetAViewers godoc
//
// @Summary    閲覧人数取得 API
// @Description
// @Tags       Viewer
// @Accept     json
// @Produce    json
// @Success    200     	{object}    GetTotalViewersResponse
// @Failure    400     	{object}    entity.ErrorResponse
// @Failure    404     	{object}    entity.ErrorResponse
// @Router     /viewers [get]
func (c ViewerController) GetTotalViewers(ctx *gin.Context) (interface{}, error) {
	viewers, err := c.ViewerUseCase.GetTotalViewers(ctx)
	if err != nil {
		return nil, err
	}
	return GetTotalViewersResponse{Viewers: viewers}, nil
}
