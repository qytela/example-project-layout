package logger

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// MakeLogEntry method for logging (like fmt but using logrus)
func MakeLogEntry(c echo.Context) *logrus.Entry {
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors:   false,
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	}

	if c == nil {
		return log.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(logrus.Fields{
		"at":         time.Now().Format("2006-01-02 15:04:05"),
		"method":     c.Request().Method,
		"uri":        c.Request().URL.String(),
		"ip":         c.Request().RemoteAddr,
		"user_agent": c.Request().UserAgent(),
	})
}
