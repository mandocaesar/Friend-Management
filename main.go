package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Friend-Management/shared"
	"github.com/Friend-Management/shared/config"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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

	//load router
	r := shared.SetupRouter(&configuration)

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
