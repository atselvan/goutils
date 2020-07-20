package utils

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	UsernameKey             = "username"
	PasswordKey             = "password"
	basicAuthRequiredErrMsg = "401 unauthorized: Basic authentication is required"
)

// BasicAuthRequired is a gin middleware for checking if basic authentication is provided in the request
// The method writes the basic auth to the gin context
// The method returns an error if basic authentication is not set
func BasicAuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.Request.Header.Get("Authorization") == "" {
			ctx.IndentedJSON(http.StatusUnauthorized, ErrResponse{Error: basicAuthRequiredErrMsg})
			log := LogFormatter{Request: ctx.Request, StatusCode: http.StatusUnauthorized, Msg: basicAuthRequiredErrMsg}
			log.Info().Println(log.Out)
			ctx.Abort()
			return
		}

		auth := strings.SplitN(ctx.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			ctx.IndentedJSON(http.StatusUnauthorized, ErrResponse{Error: basicAuthRequiredErrMsg})
			log := LogFormatter{Request: ctx.Request, StatusCode: http.StatusUnauthorized, Msg: basicAuthRequiredErrMsg}
			log.Info().Println(log.Out)
			ctx.Abort()
			return
		}

		dAuth, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, ErrResponse{Error: basicAuthRequiredErrMsg})
			log := LogFormatter{Request: ctx.Request, StatusCode: http.StatusUnauthorized, Msg: basicAuthRequiredErrMsg}
			log.Info().Println(log.Out)
			ctx.Abort()
			return
		}

		cred := strings.SplitN(string(dAuth), ":", 2)

		if len(cred) != 2 {
			ctx.IndentedJSON(http.StatusUnauthorized, ErrResponse{Error: basicAuthRequiredErrMsg})
			log := LogFormatter{Request: ctx.Request, StatusCode: http.StatusUnauthorized, Msg: basicAuthRequiredErrMsg}
			log.Info().Println(log.Out)
			ctx.Abort()
			return
		}

		ctx.Set(UsernameKey, cred[0])
		ctx.Set(PasswordKey, cred[1])

		ctx.Next()
	}
}
