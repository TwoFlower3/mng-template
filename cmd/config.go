package main

import (
	"github.com/urfave/cli"
)

const (
	FLAG_DEBUG            = "debug"
	FLAG_TEXTLOG          = "textlog"
	FLAG_TRACE            = "trace"
	FLAG_PORT             = "port"
	FLAG_HOST             = "host"
	FLAG_PG_HOST          = "pg-host"
	FLAG_PG_PORT          = "pg-port"
	FLAG_PG_SSL_MODE      = "pg-ssl-mode"
	FLAG_PG_USER          = "pg-user"
	FLAG_PG_PASSWORD      = "pg-password"
	FLAG_PG_DATABASE      = "pg-database"
	FLAG_PG_MAX_IDLE_CONN = "pg-max-idle-conn"
)

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "DEBUG",
		Name:   FLAG_DEBUG,
		Usage:  "start the server in debug mode",
	},
	cli.BoolFlag{
		EnvVar: "TEXTLOG",
		Name:   FLAG_TEXTLOG,
		Usage:  "output log in text format",
	},
	cli.BoolFlag{
		EnvVar: "TRACE",
		Name:   FLAG_TRACE,
		Usage:  "enable trace in output log",
	},
	cli.StringFlag{
		EnvVar: "HOST",
		Name:   FLAG_HOST,
		Value:  "",
		Usage:  "Server address",
	},
	cli.StringFlag{
		EnvVar: "PORT",
		Name:   FLAG_PORT,
		Value:  "8080",
		Usage:  "Server port",
	},
	cli.StringFlag{
		EnvVar: "PG_HOST",
		Name:   FLAG_PG_HOST,
		Value:  "postgresql",
		Usage:  "PostgreSQL server address",
	},
	cli.IntFlag{
		EnvVar: "PG_PORT",
		Name:   FLAG_PG_PORT,
		Value:  5432,
		Usage:  "PostgreSQL server port",
	},
	cli.StringFlag{
		EnvVar: "PG_SSL_MODE",
		Name:   FLAG_PG_SSL_MODE,
		Value:  "disable",
		Usage:  "PostgreSQL SSL Mode (disable, allow, prefer, require, verify-ca, verify-full). More: https://www.postgresql.org/docs/9.1/libpq-ssl.html",
	},
	cli.StringFlag{
		EnvVar: "PG_USER",
		Name:   FLAG_PG_USER,
		Value:  "postgres",
		Usage:  "PostgreSQL user",
	},
	cli.StringFlag{
		EnvVar: "PG_PASSWORD",
		Name:   FLAG_PG_PASSWORD,
		Value:  "strong",
		Usage:  "PostgreSQL user password",
	},
	cli.StringFlag{
		EnvVar: "PG_DATABASE",
		Name:   FLAG_PG_DATABASE,
		Value:  "strong",
		Usage:  "PostgreSQL database",
	},
	cli.IntFlag{
		EnvVar: "PG_MAX_IDLE_CONN",
		Name:   FLAG_PG_MAX_IDLE_CONN,
		Value:  5,
		Usage:  "Maximum number of concurrently idle connections for PostgreSQL",
	},
}
