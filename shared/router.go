package shared

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"

	registration "github.com/Friend-Management/module/registration"
	"github.com/Friend-Management/shared/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	ginglog "github.com/szuecs/gin-glog"
)

var (
	db                 *gorm.DB
	registerController registration.Controller
	registerService    registration.Service
)

func controllerRegistration(Db *gorm.DB) {
	registerService, err := registration.NewService(Db)

	if err != nil {
		glog.Fatalf("Failed to instantiate new Registration Service: %s", err)
		panic(fmt.Errorf("Fatal error: %s", err))
	}

	registerController := registration.NewController(registerService)
	if registerController == nil {
		glog.Fatalf("Failed to instantiate new Registration Controller")
		panic(fmt.Errorf("Fatal error: %s", "registration controller failed to instantiate"))
	}

}

//SetupRouter : function that return registered end point
func SetupRouter(configuration *config.Configuration, Db *gorm.DB) *gin.Engine {
	router := gin.New()
	duration := time.Duration(configuration.Server.LogDuration) * time.Second
	db = Db

	//register all controller
	controllerRegistration(db)

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
		registration.POST("/", registerController.RegisterUser)
	}

	return router
}
