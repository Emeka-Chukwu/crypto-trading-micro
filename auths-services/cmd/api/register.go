package api

import (
	"auths-services/data"
	"auths-services/util"
	"database/sql"
	"net/http"

	"github.com/opensaucerer/barf"
)

func (s *Server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	payload, err := util.GetBody[data.AuthsModel](r)
	authModel, err := s.repo.FindAuthsByEmail(payload.Email)
	if (err != nil) && err != sql.ErrNoRows {
		if authModel.Email != "" {
			barf.Response(w).Status(http.StatusConflict).JSON(barf.Res{Status: false, Data: nil, Message: "email already exist"})
			return
		}
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
		return
	}

	dbResp, err := s.repo.RegisterAuths(payload)
	if err != nil {
		barf.Response(w).Status(http.StatusInternalServerError).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
		return
	}

	jwTMaker, err := data.NewJWTMaker(s.config.TokenSymmetricKey)
	if err != nil {
		barf.Response(w).Status(http.StatusInternalServerError).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
		return
	}
	token, authPayload, err := jwTMaker.CreateToken(authModel.ID, s.config.AccessTokenDuration)
	if err != nil {
		barf.Response(w).Status(http.StatusInternalServerError).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
		return
	}
	var resp map[string]interface{} = map[string]interface{}{
		"token":      token,
		"user":       dbResp,
		"expires_at": authPayload.ExpiredAt,
	}
	barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
		Status:  true,
		Data:    resp,
		Message: "Register successfuly",
	})
}
