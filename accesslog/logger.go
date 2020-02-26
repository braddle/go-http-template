package accesslog

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wl := NewResponseLogger(w)
		next.ServeHTTP(wl, r)

		logrus.WithFields(
			logrus.Fields{
				"method":     r.Method,
				"request":    r.URL.Path,
				"status":     wl.Status(),
				"size":       wl.Size(),
				"user_agent": r.UserAgent(),

				//"time_local":      "",
				//"remote_addr":     "",
				//"remote_user":     "",
				//"body_bytes_sent": "",
				//"request_time":    "",
				//"http_referrer":   "",
				//"http_user_agent": "",
			}).Info()
	})
}
