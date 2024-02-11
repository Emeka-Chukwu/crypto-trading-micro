package api

import (
	"auths-services/data"
	"auths-services/util"
	"database/sql"
	"os"

	"github.com/opensaucerer/barf"
)

type Server struct {
	conn   *sql.DB
	port   string
	repo   data.AuthsRepo
	config util.Config
}

func NewServer(conn *sql.DB, config util.Config) *Server {
	repo := data.NewAuthsRepo(conn)
	return &Server{conn: conn, port: "7050", repo: repo, config: config}
}

func (s *Server) Serve() {
	if err := barf.Stark(barf.Augment{
		Port:     s.port,
		Logging:  barf.Allow(),
		Recovery: barf.Allow(),
	}); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}
	subRoutes := barf.RetroFrame("/api/v1")
	subRoutes.Post("/auths/register", s.RegisterUser, util.InputValidatorMiddleware[data.AuthsModel])
	subRoutes.Post("/auths/login", s.LoginUser, util.InputValidatorMiddleware[data.AuthsModel])

	if err := barf.Beck(); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}
}
