package logger

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type RFC3164Formatter struct{}

func (f *RFC3164Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	fmt.Fprintf(
		b, "<%d>1 %s %s %s - - - %s\n",
		entry.Level,
		entry.Time.Format(time.RFC3339),
		entry.Data["hostname"],
		entry.Data["appname"],
		entry.Message,
	)
	return b.Bytes(), nil
}

func NewLogger() *logrus.Logger {
	logger := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(RFC3164Formatter),
		Level:     logrus.InfoLevel,
	}

	return logger
}
