package routes

import (
	"github.com/ernstvorsteveld/mta-parser-service/product"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	productsRegister(router)
}

func productsRegister(router *gin.RouterGroup) {
	router.GET("/", productsHandler)
}

func productsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"products": product.GetProducts(),
	})
}
