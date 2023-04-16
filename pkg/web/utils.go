package web

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/apperror"
)

func RequiredParam(c *gin.Context, param string) (string, error) {
	val, ok := c.Params.Get(param)
	if !ok {
		err := fmt.Errorf("missing id params")
		err = apperror.WrapError(err, apperror.ErrMissingRequiredParams)
		return "", err
	}

	return val, nil
}

func AuthorizationRequired(c *gin.Context) (string, error) {
	tokenValue := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(tokenValue) != 2 {
		err := fmt.Errorf("invalid token")
		err = apperror.WrapError(err, apperror.ErrUnauthorized)
		return "", err
	}
	token := tokenValue[1]

	return token, nil
}
