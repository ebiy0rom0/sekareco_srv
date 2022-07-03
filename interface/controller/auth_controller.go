package controller

import (
	"fmt"
	"net/http"
	"sekareco_srv/infra/logger"
	"sekareco_srv/interface/database"
	"sekareco_srv/logic/auth"
	"time"
)

type AuthController struct {
	logic auth.AuthLogic
}

func NewAuthController(h database.SqlHandler) *AuthController {
	return &AuthController{
		logic: auth.AuthLogic{
			Repository: &database.PersonRepository{
				Handler: h,
			},
		},
	}
}

// using middleware
func (c *AuthController) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := c.logic.GetInHeaderToken(r)
		if len(token) == 0 {
			logger.Logger.Warn(fmt.Errorf("%s", "unauthorized"))
			w.Header().Set(auth.RESPONSE_HEADER, auth.MESSAGE_UNAUTHORIZED)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: person ID getting from request parameter
		pid := 1
		if !c.logic.IsEnabledToken(pid, token) {
			logger.Logger.Warn(fmt.Errorf("%s", "invalid token"))
			w.Header().Set(auth.RESPONSE_HEADER, auth.MESSAGE_INVALID_TOKEN)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set(auth.RESPONSE_HEADER, auth.MESSAGE_OK)
		next.ServeHTTP(w, r)
	})
}

func (c *AuthController) DeleteExpiredToken() {
	// t := time.NewTicker(auth.EXPIRED_TOKEN_DELETE_SPAN)
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		<-t.C
		c.logic.DeleteExpiredToken()
	}
}
