package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/Friend-Management/shared/config"
	"github.com/Friend-Management/shared/data"
	"github.com/Friend-Management/module/registration"

	"github.com/golang/glog"
	"github.com/stretchr/testify/assert"
	"github.com/Friend-Management/shared"
)

func TestServiceRegisterUser(t *testing.T){
	cfg, err := config.New()

	if err != nil {
		glog.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error on load configuration : %s ", err))
	}

	db, error := data.NewDbFactory(cfg)
	if error != nil {
		glog.Fatalf("Failed to create new database Instance: %s", err)
	}

	database, error := db.DBConnection()
	if error != nil {
		glog.Fatalf("Failed to create connection: %s", err)
	}

	userservice , error:= registration.NewService(database)

	result, err, id := userservice.CreateUser("test@test.com")

	assert.NotEqual(t, id, 0)
	assert.True(t, result)
	assert.Nil(t, err)


}

func TestRegisterUser(t *testing.T) {

	cfg, err := config.New()

	if err != nil {
		glog.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error on load configuration : %s ", err))
	}

	assert.Empty(t, err)

	payload := bytes.NewBuffer([]byte(`{"email":"test@user.com"}`))

	db, error := data.NewDbFactory(cfg)
	if error != nil {
		glog.Fatalf("Failed to create new database Instance: %s", err)
	}

	database, error := db.DBConnection()
	if error != nil {
		glog.Fatalf("Failed to create connection: %s", err)
	}

	routerInstance := shared.NewRouter(cfg, database)

	router := routerInstance.SetupRouter()

	//Perform http request
	response := DispatchRequest(router, "POST", "/v1/api/register", payload)

	assert.Equal(t, http.StatusTemporaryRedirect, response.Code)
}
