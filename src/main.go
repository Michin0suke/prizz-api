package main

import (
	"github.com/Michin0suke/prizz-api/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/")
	{
		api.GET("deadline", controller.DeadlineGET)
		api.GET("new", controller.NewGET)
		api.GET("winner", controller.WinnerGET)
		api.GET("category/:category", controller.CategoryGET)
		api.GET("search/:id", controller.SearchGET)
		api.GET("way/twitter", controller.TwitterGET)
		api.GET("total_number", controller.TotalNumberGET)
	}
	router.Run(":9000")
}
