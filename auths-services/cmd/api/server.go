package api

import (
	"auths-services/data"
	"auths-services/util"
	"database/sql"
	"net/http"
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
	barf.Get("/", func(w http.ResponseWriter, r *http.Request) {

		barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
			Status:  true,
			Data:    nil,
			Message: "Hello World",
		})
	})
	subRoutes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
			Status:  true,
			Data:    nil,
			Message: "Hello World",
		})
	})

	if err := barf.Beck(); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}
}
