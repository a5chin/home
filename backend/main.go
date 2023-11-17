package main

import (
	"fmt"
	"log"
	"net/http"

	"home/config"
	"home/docs"
	"home/entity"
	"home/infrastructure/driver"
	"home/infrastructure/middleware"

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

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction(db))

	app.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "It works")
	})

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
