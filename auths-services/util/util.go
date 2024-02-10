package util

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBody[T any](r *http.Request) T {
	var payload T
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return payload
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return payload
	}
	return payload
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
