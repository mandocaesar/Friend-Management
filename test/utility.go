package test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/Friend-Management/shared/config"
	"github.com/golang/glog"
)

//DispatchRequest : function to mock http-request for testing purpose
func DispatchRequest(req http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, body)
	request.Header.Set("Content-Type", "application/json")

	glog.V(2).Info(request)

	response := httptest.NewRecorder()
	req.ServeHTTP(response, request)

	return response
}

//LoadConfiguration : load test configuration
func LoadConfiguration() (*config.Configuration, error) {
	cfg, err := config.New()

	return cfg, err
}
