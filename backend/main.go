package main

import (
	"fmt"
	"log"
	"net/http"

	"home/config"
	"home/controller"
	"home/docs"
	"home/entity"
	"home/infrastructure/driver"
	"home/infrastructure/middleware"
	"home/infrastructure/repository"
	"home/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title		home
// @description	home
// @version		1.0
func main() {
	// Load config
	conf := config.Load()
	db := driver.NewDB(conf)

	// Dependency Injection
	articleRepository := repository.NewArticleRepository()
	lpRepository := repository.NewLPRepository()

	articleUseCase := usecase.NewArticleUseCase(articleRepository)
	lpUseCase := usecase.NewLPUseCase(lpRepository)

	articleController := controller.NewArticleController(articleUseCase)
	lpController := controller.NewLPController(lpUseCase)

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction(db))

	app.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "It works")
	})

	{
		v1 := app.Group("/api/v1")

		articleRouter := v1.Group("articles")
		articleRouter.GET("", handleResponse(articleController.GetArticles))

		lpRouter := v1.Group("viewers")
		lpRouter.GET("", handleResponse(lpController.GetTotalViewers))
	}

	runApp(app, conf)
}

func runApp(app *gin.Engine, conf *config.Config) {
	if config.ExistEnvFile() {
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", conf.PORT)
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http"}

		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		log.Println(fmt.Sprintf("http://localhost:%s", conf.PORT))
		log.Println(fmt.Sprintf("http://localhost:%s/swagger/index.html", conf.PORT))
	}
	app.Run(fmt.Sprintf("%s:%s", conf.HOSTNAME, conf.PORT))
}

func handleResponse(f func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := f(c)
		if err != nil {
			e, ok := err.(*entity.Error)
			if ok {
				c.JSON(e.Code, entity.ErrorResponse{Message: err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
			}
			c.Abort()
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
