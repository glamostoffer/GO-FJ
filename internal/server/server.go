package server

import (
	"GO-FJ/internal/config"
	conn "GO-FJ/pkg/postgres_connector"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg          *config.Config
	serverEngine *gin.Engine
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg:          cfg,
		serverEngine: gin.New(),
	}
}

func (s Server) Run() error {

	psqlConn, err := conn.NewConnector(s.cfg.Postgres)
	if err != nil {
		return err
	}
	fmt.Printf(`Connected to DB with connector %s %s`, psqlConn, "\n")
	// userUsecase := userUsecase.New(msqlConn)
	// userGroup := s.serverEngine.Group("/user")
	// userHandlers := userHandlers.New(userUsecase)
	// routeUserHandlers(userGroup, userHandlers)

	err = s.serverEngine.Run(fmt.Sprintf(":%s", s.cfg.HTTP.Port))

	return err
}
