package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/Friend-Management/shared"
	"github.com/Friend-Management/shared/config"
	"github.com/golang/glog"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {

	cfg, err := config.New()

	if err != nil {
		glog.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error on load configuration : %s ", err))
	}

	assert.Empty(t, err)

	payload := bytes.NewBuffer([]byte(`{"email":"test@user.com"}`))

	router := shared.SetupRouter(cfg)

	//Perform http request
	response := DispatchRequest(router, "POST", "/v1/api/register", payload)

	assert.Equal(t, http.StatusTemporaryRedirect, response.Code)
}
