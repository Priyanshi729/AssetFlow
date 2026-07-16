package middleware

import (
	"AssetFlow/models"
	"AssetFlow/utils"
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/form3tech-oss/jwt-go"
)

type ContextKeys string

const (
	userContext ContextKeys = "__userContext"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("x-api-key")
		if tokenString == "" {
			utils.RespondError(w, http.StatusUnauthorized, nil, "token header missing")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized, err, "invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.RespondError(w, http.StatusUnauthorized, nil, "invalid token")
			return
		}

		user := &models.User{
			UserID: claims["userId"].(string),
			Role:   claims["role"].(string),
		}

		ctx := context.WithValue(r.Context(), userContext, user)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func UserContext(r *http.Request) *models.User {
	if user, ok := r.Context().Value(userContext).(*models.User); ok {
		return user
	}
	return nil
}

func RequireRoles(next http.Handler, roles ...string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := UserContext(r)

		if user == nil {
			utils.RespondError(w, http.StatusUnauthorized, nil, "Unauthorized")
			return
		}

		for _, role := range roles {
			if user.Role == role {
				next.ServeHTTP(w, r)
				return
			}
		}

		utils.RespondError(w, http.StatusForbidden, nil, "Access Denied")
	})
}
