package main

import (
	"github.com/search-api/docs"
	"github.com/search-api/elasticConn"
	"github.com/search-api/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {

	// release mode
	gin.SetMode(gin.ReleaseMode)

	// elasticsearch connect
	elasticConn.InitEsClient()

	// swagger setting
	docs.SwaggerInfo.Title = "SEARCH-API Swagger"
	//docs.SwaggerInfo.Description = "search-api PRJ swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""

	e := setRouter()
	e.Run()
}

func setRouter() *gin.Engine {
	/*
		Router 설정
		단일 처리
		e := gin.Default()
		c := controller.NewController()
		e.GET("/search", c.ProdSearch)
	*/
	e := gin.New()
	rt := router.NewRouter()

	//prod router register
	rt.ProdRegister(e.Group("/prod"))

	// swagger
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return e
}
