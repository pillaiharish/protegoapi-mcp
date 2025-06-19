package middleware

import (
	"fmt"
	"context"
	"net/http"
	"strings"
	"time"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

const  secret = "I_AM_GROOT"

type ctxKey string

var jwtKey ctxKey = "jwt"

func RequireJWT(next http.Handler) http.Handler{
	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println("Authorization header:", authHeader)

		raw := strings.TrimPrefix( authHeader, "Bearer ")
		fmt.Println("Raw token:", raw)

		if raw == "" {
			http.Error(w, "missing token",http.StatusUnauthorized)
			return
		}
		tok, err := jwt.ParseString(raw, jwt.WithVerify(false))
		if err!=nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		exp, hasExp := tok.Expiration()
		if hasExp && exp.Before(time.Now()) {
			http.Error(w, "expired token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(),jwtKey ,tok)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

