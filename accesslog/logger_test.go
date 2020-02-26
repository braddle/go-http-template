package accesslog_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/braddle/go-http-template/accesslog"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type LoggerSuite struct {
	suite.Suite
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}

const (
	Content   = "This is the repose body"
	url       = "/people/123456"
	UserAgent = "Braddle_Client/0.1"
)

func (s *LoggerSuite) TestAccessLogging() {
	logBuf := bytes.NewBufferString("")
	logrus.SetOutput(logBuf)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	r := mux.NewRouter()
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(Content))
	})
	r.Use(accesslog.Logger)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", UserAgent)
	r.ServeHTTP(httptest.NewRecorder(), req)

	access := make(map[string]interface{})
	sc := bufio.NewScanner(logBuf)
	sc.Scan()

	json.Unmarshal(sc.Bytes(), &access)

	// Request
	s.Equal(url, access["request"].(string))
	s.Equal(http.MethodGet, access["method"].(string))
	s.Equal(UserAgent, access["user_agent"].(string))

	// Response
	s.Equal(http.StatusTeapot, int(access["status"].(float64)))
	s.Equal(len(Content), int(access["size"].(float64)))
}
