// Package logger
package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type LogHandler struct {
	logoutFile *os.File
	*slog.Logger
}

func NewLogHandler() *LogHandler {
	l := &LogHandler{}
	logfile, err := os.OpenFile("log.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Printf("error creating Log File")
		return nil
	}
	l.logoutFile = logfile
	writer := io.MultiWriter(os.Stdout, logfile)
	l.Logger = slog.New(slog.NewJSONHandler(writer, nil))
	l.Info("logHandler created")

	return l
}

func (l *LogHandler) Close() error {
	err := l.logoutFile.Close()
	if err != nil {
		return err
	}
	return nil
}

// logger.Info("hello, world", "user", os.Getenv("USER"))
