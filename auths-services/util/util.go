package util

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBody[T any](r *http.Request) (T, error) {
	var payload T
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return payload, err
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return payload, err
	}
	return payload, err
}

func GetUrlParams[T any](c *gin.Context) T {
	var payload T
	err := c.ShouldBindUri(&payload)
	if err != nil {
		return payload
	}
	return payload
}

func GetUrlQueryParams[T any](c *gin.Context) T {
	var payload T
	err := c.ShouldBindQuery(&payload)
	if err != nil {
		return payload
	}
	return payload
}
