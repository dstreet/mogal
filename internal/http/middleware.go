package http

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/dstreet/mogal/internal/user"
)

type AuthMiddleware struct {
	UserRepository user.UserRepository
	TokenProvider  TokenProvider
	Logger         *slog.Logger
}

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{"user"}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		split := strings.Split(authorization, "Bearer ")
		if len(split) < 2 {
			m.Logger.Info("no bearer token provided")
			next.ServeHTTP(w, r)
			return
		}

		token := split[1]

		userID, err := m.TokenProvider.VerifyToken(token)
		if err != nil {
			m.Logger.Warn("failed to verify token", "err", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := m.UserRepository.GetWithID(r.Context(), userID)
		if err != nil {
			m.Logger.Error("failed to get user", "ID", userID)
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, &user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func UserForContext(ctx context.Context) *user.User {
	u, _ := ctx.Value(UserCtxKey).(*user.User)
	return u
}
