package controller

import (
	"errors"
	"home/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	ArticleUseCase
}

func NewArticleController(u ArticleUseCase) *ArticleController {
	return &ArticleController{ArticleUseCase: u}
}

type GetArticlesQuery struct {
	Limit  *int `form:"limit,omitempty"`
	Offset *int `form:"offset,omitempty"`
}

type GetArticlesResponse struct {
	Articles []*entity.Article `json:"articles"`
}

// GetArticles godoc
//
// @Summary    投稿一覧取得API
// @Description
// @Tags       Post
// @Accept     json
// @Produce    json
// @Param      limit   	query       string  false   "limit"
// @Param      offset  	query       string  false   "offset"
// @Success    200     	{object}    GetArticlesResponse
// @Failure    400     	{object}    entity.ErrorResponse
// @Failure    404     	{object}    entity.ErrorResponse
// @Router     /articles [get]
func (c ArticleController) GetArticles(ctx *gin.Context) (interface{}, error) {
	var query GetArticlesQuery
	err := ctx.ShouldBind(&query)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	if query.Limit == nil && query.Offset != nil {
		return nil, entity.WrapError(http.StatusBadRequest, errors.New("can't use offset without limit"))
	}
	articles, err := c.ArticleUseCase.GetArticles(ctx, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}
	return GetArticlesResponse{Articles: articles}, nil
}
