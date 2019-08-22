package middleware

import (
	"github.com/labstack/echo"
	"time"
	"net"
	"net/http"
	log "github.com/sirupsen/logrus"
	"fmt"
)

// Logger function for writing all request log into console
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := "acticity-service"
		l := log.StandardLogger()

		start := time.Now()
		req := c.Request()
		res := c.Response()

		remoteAddr := req.RemoteAddr
		if ip := req.Header.Get(echo.HeaderXRealIP); ip != "" {
			remoteAddr = ip
		} else if ip = req.Header.Get(echo.HeaderXForwardedFor); ip != "" {
			remoteAddr = ip
		} else {
			remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
		}

		entry := l.WithFields(log.Fields{
			"request": req.RequestURI,
			"method":  req.Method,
			"remote":  remoteAddr,
		})

		if reqID := req.Header.Get("X-Request-Id"); reqID != "" {
			entry = entry.WithField("request_id", reqID)
		}

		entry.Info("started handling request")

		if err := next(c); err != nil {
			c.Error(err)
		}

		latency := time.Since(start)

		entry.WithFields(log.Fields{
			"size":        res.Size,
			"status":      res.Status,
			"text_status": http.StatusText(res.Status),
			"took":        latency,
			fmt.Sprintf("#%s.latency", name): latency.Nanoseconds(),
		}).Info("completed handling request")

		return nil
	}
}