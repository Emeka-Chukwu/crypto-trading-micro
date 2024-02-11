package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/opensaucerer/barf"
)

func ValidateInput(payload interface{}) (map[string]string, error) {
	v := validator.New()
	err := v.Struct(payload)

	errors := make(map[string]string)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {

			errors[strings.ToLower(e.Field())] = e.ActualTag()
		}
	}
	return errors, err
}

func ValidatorMiddleware[T any](w http.ResponseWriter, r *http.Request) {
	var payload T
	body, err := io.ReadAll(r.Body)
	if err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Data:    nil,
			Message: "error",
		})
		return
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Data:    nil,
			Message: "error",
		})
		return
	}

	_, err = ValidateInput(payload)
	if err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Data:    nil,
			Message: "error",
		})
		return
	}
	// c.Set("body", payload)

}

func InputValidatorMiddleware[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload T
		////
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r.Body)
		if err != nil {
			barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
			return
		}
		body := buf.Bytes()
		r.Body = io.NopCloser(&buf)
		//////
		if err != nil {
			barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
			return
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
			return
		}
		_, err = ValidateInput(payload)
		if err != nil {
			barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{Status: false, Data: err.Error(), Message: "error"})
			return
		}
		next.ServeHTTP(w, r)

	})
}
