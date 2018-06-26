package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Friend-Management/config"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	ginglog "github.com/szuecs/gin-glog"
)

var (
	runMigration  bool
	runSeeder     bool
	configuration config.Configuration
)

func init() {
	flag.BoolVar(&runMigration, "migrate", false, "run db migration")
	flag.Parse()

	glog.V(2).Info("Initilizing server...")
	cfg, err := config.New()

	if err != nil {
		glog.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error on load configuration : %s ", err))
	}
	configuration = *cfg

}

func main() {
	glog.V(2).Infof("Server run on mode: %s", configuration.Server.Mode)
	gin.SetMode(configuration.Server.Mode)
	r := setupRouter()

	srv := &http.Server{
		Addr:    configuration.Server.Addr,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		glog.Fatalf("Failed to start server: %s", err)
		panic(fmt.Errorf("Fatal error failed to start server: %s", err))
	}

	glog.Info("Shutting down server")
	// give n seconds for server to shutdown gracefully
	duration := time.Duration(configuration.Server.ShutdownTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Fatalf("Failed to shut down server gracefully: %s", err)
	}
	glog.Info("Server shutted down")
}

func setupRouter() *gin.Engine {
	router := gin.New()
	duration := time.Duration(configuration.Server.LogDuration) * time.Second

	//middleware setup
	router.Use(ginglog.Logger(duration), gin.Recovery())

	// diagnostic endpoint
	web := router.Group("/api")
	{
		web.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message":    "OK",
				"serverTime": time.Now().UTC(),
				"version":    "1.0.0",
			})
		})
	}

	return router
}
