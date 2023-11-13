package main

import (
	"GO-FJ/internal/config"
	"GO-FJ/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Error reading config: %s", err)
		return
	}

	err = server.New(cfg).Run()
	if err != nil {
		logrus.Fatalf(err.Error())
		return
	}
}
