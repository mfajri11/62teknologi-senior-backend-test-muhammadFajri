package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/apperror"
	log "github.com/rs/zerolog/log"
)

type ErrorAPI struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Err ErrorAPI `json:"error"`
}

func WriteError(c *gin.Context, err error) {
	var apiErr apperror.APIError
	if errors.As(err, &apiErr) {
		statusCode, code, desc := apiErr.Message()
		log.Err(err).Msg(code)
		c.JSON(statusCode, ErrorResponse{ErrorAPI{Code: code, Description: desc}})
		return
	}

	// if not apperror.APIError it considered as internal error with no body
	c.Status(http.StatusInternalServerError)
}
