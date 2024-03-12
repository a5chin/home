package controller

import (
	"github.com/gin-gonic/gin"
)

type LPController struct {
	LPUseCase
}

func NewLPController(u LPUseCase) *LPController {
	return &LPController{u}
}

type GetTotalViewersResponse struct {
	Viewers uint `json:"viewers"`
}

// GetAViewers godoc
//
// @Summary    閲覧人数取得 API
// @Description
// @Tags       LP
// @Accept     json
// @Produce    json
// @Success    200     	{object}    GetTotalViewersResponse
// @Failure    400     	{object}    entity.ErrorResponse
// @Failure    404     	{object}    entity.ErrorResponse
// @Router     /viewers [get]
func (c LPController) GetTotalViewers(ctx *gin.Context) (interface{}, error) {
	viewers, err := c.LPUseCase.GetTotalViewers(ctx)
	if err != nil {
		return nil, err
	}
	return GetTotalViewersResponse{Viewers: viewers}, nil
}
