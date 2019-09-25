package main

import (
	"net/http"

	"github.com/Michin0suke/prizz-api/src/controller"
	"github.com/Michin0suke/prizz-api/src/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	flags := util.GetFlags()
	if *flags.Mode != "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	store := cookie.NewStore([]byte("super-ebimayo"))

	router := gin.Default()
	router.Use(sessions.Sessions("oauth-twitter-settion", store))

	api := router.Group("/")
	{
		api.GET("contents", controller.ContentsGET)
		api.GET("search/:id", controller.SearchGET)
		api.GET("total_number", controller.TotalNumberGET)
		api.GET("/login/twitter", controller.TwitterLogin)
		api.GET("/login/twitter/callback", controller.TwitterCallback)
		api.GET("/login/twitter/is_logged_in", controller.TwitterIsLoggedIn)
		api.GET("/twitter/follow/:user_param", controller.TwitterFollow)
		api.GET("/twitter/retweet/:id", controller.TwitterRetweet)
		api.GET("/twitter/favorite/:id", controller.TwitterFavorite)
		api.POST("/twitter/reply/:id", controller.TwitterReply)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"error": "Endpoint is not valid."})
	})

	switch *flags.Mode {
	case "production":
		router.Run(":9000")
	case "development":
		router.Run(":9999")
	default:
		router.Run(":9000")
	}
}
