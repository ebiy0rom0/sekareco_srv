package controller

import (
	"fmt"
	"net/http"
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"
	"time"
)

type AuthController struct {
	authLogic model.AuthLogic
}

func NewAuthController(a model.AuthLogic) *AuthController {
	return &AuthController{
		authLogic: a,
	}
}

// using middleware
func (c *AuthController) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := c.authLogic.GetHeaderToken(r)
		if len(token) == 0 {
			logger.Logger.Warn(fmt.Errorf("%s", "unauthorized"))
			w.Header().Set(model.RESPONSE_HEADER, model.HEADER_UNAUTHORIZED)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: person ID getting from request parameter
		pid := 1
		if !c.authLogic.IsEnabledToken(pid, token) {
			logger.Logger.Warn(fmt.Errorf("%s", "invalid token"))
			w.Header().Set(model.RESPONSE_HEADER, model.HEADER_INVALID_TOKEN)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set(model.RESPONSE_HEADER, model.HEADER_DONE)
		next.ServeHTTP(w, r)
	})
}

func (c *AuthController) DeleteExpiredToken() {
	// t := time.NewTicker(auth.EXPIRED_TOKEN_DELETE_SPAN)
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		<-t.C
		c.authLogic.DeleteExpiredToken()
	}
}
