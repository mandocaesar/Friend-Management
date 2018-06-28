package shared

import (
	"net/http"
	"time"

	"github.com/Friend-Management/shared/config"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
)

//SetupRouter : function that return registered end point
func SetupRouter(configuration *config.Configuration) *gin.Engine {
	router := gin.New()
	duration := time.Duration(configuration.Server.LogDuration) * time.Second

	//middleware setup
	router.Use(ginglog.Logger(duration), gin.Recovery())

	// diagnostic endpoint
	web := router.Group("v1/api")
	{
		web.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message":    "OK",
				"serverTime": time.Now().UTC(),
				"version":    "1.0.0",
			})
		})
	}

	registration := router.Group("v1/api/register")
	{
		registration.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message":    "OK",
				"serverTime": time.Now().UTC(),
				"version":    "1.0.0",
			})
		})
		registration.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"message":    "OK",
				"serverTime": time.Now().UTC(),
				"version":    "1.0.0",
			})
		})
	}

	return router
}
