package test

import (
	"io"
	"net/http"
	"net/http/httptest"
)

//DispatchRequest : function to mock http-request for testing purpose
func DispatchRequest(req http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, body)
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	req.ServeHTTP(response, request)

	return response
}
