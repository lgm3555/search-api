package router

import (
	"github.com/gin-gonic/gin"

	controller "github.com/search-api/controller"
)

func (r *Router) ProdRegister(router *gin.RouterGroup) {
	c := controller.NewController()

	router.GET("/", c.GetProd)
	router.POST("/", c.PostProd)
}
