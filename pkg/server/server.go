package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/TwoFlower3/mng-template/pkg/logger"
	"github.com/TwoFlower3/mng-template/pkg/models"
	"github.com/TwoFlower3/mng-template/pkg/types"
	"github.com/TwoFlower3/mng-template/pkg/utils"
)

// Server dummy
type Server struct {
	log        *logger.Logger
	httpServer *http.Server
	db         *sqlx.DB

	DB models.Datastore
}

// Options dummy
type Options struct {
	Logger       *logger.Logger
	WriteTimeout time.Duration
	Address      string
	DB           types.DBOptions
}

// New create server
func New(options Options) *Server {
	s := &Server{
		log: options.Logger,
		httpServer: &http.Server{
			Addr: options.Address,
		},
	}
	s.connectDB(options.DB)
	return s
}

func (server *Server) registerHandler() {
	server.httpServer.Handler = createRouter(RouterOptions{
		log: server.log,
	})
}

func (server *Server) connectDB(dbOptions types.DBOptions) error {
	server.log.WithFields(logrus.Fields{
		"Host":         dbOptions.Host,
		"Port":         dbOptions.Port,
		"User":         dbOptions.User,
		"Password":     utils.HidePass(dbOptions.Password),
		"Database":     dbOptions.Database,
		"SSL Mode":     dbOptions.SSLMode,
		"MaxIdleConns": dbOptions.MaxIdleConn,
	}).Debug("Database params")

	db, err := models.CreateDB(dbOptions.Host,
		dbOptions.Port,
		dbOptions.User,
		dbOptions.Password,
		dbOptions.Database,
		dbOptions.SSLMode)
	if err != nil {
		return fmt.Errorf("connectDB error: %+v", err)
	}

	server.db = db.DB
	server.DB = db

	return nil
}

// Start server
func (server *Server) Start() error {
	server.registerHandler()
	server.log.WithField("Address", server.Address()).Info("Starting server")
	return server.httpServer.ListenAndServe()
}

// Shutdown server
func (server *Server) Shutdown(sec time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), sec*time.Second)
	defer cancel()
	return server.httpServer.Shutdown(ctx)
}

// Address of server
func (server *Server) Address() string {
	return server.httpServer.Addr
}
