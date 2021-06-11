package middleware

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/hiromoon/go-api-reference/domain/model/user"
	"net/http"
	"strings"

	"github.com/hiromoon/go-api-reference/infra"
)

type BasicAuthenticationMiddleware struct {
	redis      *infra.Redis
	repository *infra.UserRepository
}

func NewBasicAuthenticationMiddleware(redis *infra.Redis, repo *infra.UserRepository) *BasicAuthenticationMiddleware {
	return &BasicAuthenticationMiddleware{
		redis:      redis,
		repository: repo,
	}
}

func (m *BasicAuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := m.authentication(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "context-user", user)))
	})
}

func (m *BasicAuthenticationMiddleware) authentication(authzHeader string) (*user.User, error) {
	tokens := strings.Split(authzHeader, " ")
	if tokens[0] != "Basic" {
		return nil, errors.New("error")
	}

	dec, err := base64.StdEncoding.DecodeString(tokens[1])
	if err != nil {
		return nil, err
	}

	cred := strings.Split(string(dec), ":")
	userID, password := cred[0], cred[1]
	user := user.User{}
	err = m.redis.RunWithCache(userID, &user, func(dest interface{}) error {
		dest, err := m.repository.Get(userID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("authentication error")
	}
	return &user, nil
}
