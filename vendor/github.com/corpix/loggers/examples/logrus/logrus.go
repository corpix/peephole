package main

import (
	"github.com/sirupsen/logrus"

	logrusLogger "github.com/corpix/loggers/logger/logrus"
)

func main() {
	l := logrusLogger.New(logrus.New())

	_, err := l.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}

	l.Debug("Hidden")
	l.Print("Info")
	l.Error("Error")
	l.Fatal("Fatal")
}
