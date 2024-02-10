package api

import (
	"auths-services/data"
	"auths-services/util"
	"net/http"

	"github.com/opensaucerer/barf"
)

func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	payload := util.GetBody[data.AuthsModel](r)
	authModel, err := s.repo.FindAuthsByEmail(payload.Email)
	if err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Data:    nil,
			Message: "error",
		})
	}
	matched, err := s.repo.PasswordMatches(payload.Password, authModel.Password)
	if err != nil || !matched {
		barf.Response(w).Status(http.StatusUnauthorized).JSON(barf.Res{
			Status:  false,
			Data:    nil,
			Message: "error",
		})
	}

}
