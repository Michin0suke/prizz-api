package main

import (
	"github.com/Michin0suke/prizz-api/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/")
	{
		api.GET("contents", controller.ContentsGET)
		api.GET("search/:id", controller.SearchGET)
		api.GET("total_number", controller.TotalNumberGET)
	}
	router.Run(":9000")
}
