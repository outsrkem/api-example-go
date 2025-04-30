package logger

import (
	"bytes"
	"fmt"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MyFormatter struct {
}

func (MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	asctime := entry.Time.Format(time.DateTime + " -0700")
	level := entry.Level.String()
	var caller string
	if entry.Caller != nil {
		caller = fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	} else {
		caller = "?:?"
	}
	requestID := "-"
	if val, exists := entry.Data["requestID"]; exists {
		if id, ok := val.(string); ok {
			requestID = id
		}
	}

	fmt.Fprintf(b, "[%s] [%s] [%-7s] [%s] %s\n",
		asctime,
		requestID,
		level,
		caller,
		entry.Message,
	)
	return b.Bytes(), nil
}

func FromContext(c *gin.Context) *logrus.Entry {
	requestID := "-"
	if val, exists := c.Keys["RequestID"]; exists {
		if id, ok := val.(string); ok {
			requestID = id
		}
	}
	return logrus.WithFields(logrus.Fields{"requestID": requestID})
}

func NewLogrus() {
	logrus.SetFormatter(&MyFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
