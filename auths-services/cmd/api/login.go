package api

import (
	"auths-services/data"
	"auths-services/util"
	"net/http"

	"github.com/opensaucerer/barf"
)

func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	payload, err := util.GetBody[data.AuthsModel](r)
	authModel, err := s.repo.FindAuthsByEmail(payload.Email)
	if err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
		return
	}
	matched, err := s.repo.PasswordMatches(payload.Password, authModel.Password)
	if err != nil || !matched {
		barf.Response(w).Status(http.StatusUnauthorized).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
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
		"user":       authModel,
		"expires_at": authPayload.ExpiredAt,
	}
	barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
		Status:  true,
		Data:    resp,
		Message: "Login successfuly",
	})
}
