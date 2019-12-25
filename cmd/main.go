package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/TwoFlower3/mng-template/pkg/logger"
	"github.com/TwoFlower3/mng-template/pkg/server"
)

var version string
var logs *logger.Logger

const (
	defaultVersion = "v0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "template"
	app.Usage = "Template for projects"
	app.Version = version
	app.Flags = flags
	app.Commands = commands
	app.Before = setupLogger
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func setupLogger(ctx *cli.Context) error {
	logs = logger.New(version, defaultVersion)

	if ctx.Bool(FLAG_TRACE) {
		logs.EnableTrace(true)
	}

	if ctx.Bool(FLAG_DEBUG) {
		logs.SetLogMode(logrus.DebugLevel)
	} else {
		logs.SetLogMode(logrus.InfoLevel)
	}

	if ctx.Bool(FLAG_TEXTLOG) {
		logs.SetLogFormatter(logger.TextLogFormat)
	} else {
		logs.SetLogFormatter(logger.JSONLogFormat)
	}

	logs.WithFields(logrus.Fields{
		"LogLevel": logs.GetLevel(),
		"Format":   logs.Format,
	}).Info("Logger setup")

	return nil
}

func run(ctx *cli.Context) error {
	s := server.New(server.Options{
		Logger:       logs,
		WriteTimeout: 600 * time.Second,
		Address:      fmt.Sprintf("%s:%d", ctx.String(FLAG_HOST), ctx.Int(FLAG_PORT)),
		DB: server.DBOptions{
			Host:        ctx.String(FLAG_PG_HOST),
			Port:        ctx.String(FLAG_PG_PORT),
			SSLMode:     ctx.String(FLAG_PG_SSL_MODE),
			MaxIdleConn: ctx.Int(FLAG_PG_MAX_IDLE_CONN),
			Database:    ctx.String(FLAG_PG_DATABASE),
			User:        ctx.String(FLAG_PG_USER),
			Password:    ctx.String(FLAG_PG_PASSWORD),
		},
	})

	// TODO: ADD exit or error
	go s.Start()
	waitShutdownSign()

	logs.Info("Server shutdown initialized")
	return s.Shutdown(time.Second * 5)
}

func waitShutdownSign() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
