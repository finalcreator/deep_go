package simpleserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	Routes()
}

func TestRoutes(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(recorder, req)
	t.Log("status code: ", recorder.Code)
	t.Log("body: ", recorder.Body.String())
}
