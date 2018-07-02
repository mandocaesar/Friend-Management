package test

import (
	"fmt"
	"testing"

	"github.com/Friend-Management/shared/config"
	"github.com/golang/glog"
	"github.com/stretchr/testify/assert"
)

func TestReadConfigurationFromFile(t *testing.T) {
	cfg, err := config.New()

	if err != nil {
		glog.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error on load configuration : %s ", err))
	}

	assert.Empty(t, err)

	configuration := *cfg

	// test to configuration object
	assert.NotEmpty(t, configuration)

	//test config  value
	assert.Equal(t, configuration.Server.Mode, "debug")
	assert.Equal(t, configuration.Server.Addr, ":8080")
	assert.Equal(t, configuration.Server.LogDuration, 3)
	assert.Equal(t, configuration.Server.ShutdownTimeout, 5)
	assert.Equal(t, configuration.Server.BaseURL, "http://localhost:8080")

}
