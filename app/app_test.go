package app_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/braddle/go-http-template/app"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/suite"
)

type ApplicationSuite struct {
	suite.Suite
}

func TestApplicationSuite(t *testing.T) {
	suite.Run(t, new(ApplicationSuite))
}

func (s *ApplicationSuite) TestHealthCheck() {
	router := mux.NewRouter()

	app.New(router)

	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(`{"status": "OK", "errors": []}`, string(body))
}
