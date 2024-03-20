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
	return &ArticleController{u}
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
// @Summary    記事一覧取得 API
// @Description
// @Tags       Article
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

type GetArticleByIDResponse struct {
	Article *entity.Article `json:"article"`
}

// GetArticleByID godoc
//
// @Summary		記事取得 API
// @Description
// @Tags		Article
// @Accept		json
// @Produce		json
// @Param		articleId	path	string	true	"Article ID"
// @Success		200     	{object}    GetArticleByIDResponse
// @Failure		400     	{object}    entity.ErrorResponse
// @Failure		404     	{object}    entity.ErrorResponse
// @Router		/articles/{articleId} [get]
func (c ArticleController) GetArticleByID(ctx *gin.Context) (interface{}, error) {
	articleId := ctx.Param("articleId")
	article, err := c.ArticleUseCase.GetArticleByID(ctx, articleId)
	if err != nil {
		return nil, err
	}
	return GetArticleByIDResponse{Article: article}, nil
}

type GetTotalViewersResponse struct {
	TotalViewers *uint `json:"totalViewers"`
}

// GetTotalViewers godoc
//
// @Summary		全記事閲覧数 API
// @Description
// @Tags		Article
// @Accept		json
// @Produce		json
// @Success		200     	{object}    GetTotalViewersResponse
// @Failure		400     	{object}    entity.ErrorResponse
// @Failure		404     	{object}    entity.ErrorResponse
// @Router		/articles/totalViewers [get]
func (c ArticleController) GetTotalViewers(ctx *gin.Context) (interface{}, error) {
	totalViewers, err := c.ArticleUseCase.GetTotalViewers(ctx)
	if err != nil {
		return nil, err
	}
	return GetTotalViewersResponse{TotalViewers: totalViewers}, nil
}
