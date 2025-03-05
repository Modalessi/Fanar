package fanar

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func authorized(next FanaerHandler, jwtSecret string) FanaerHandler {
	return func(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			return errorResponse(401, fmt.Errorf("authorization header is missing")), nil
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := VerfiyToken(tokenString, jwtSecret)
		if err != nil {
			return errorResponse(401, fmt.Errorf("invalid token")), nil
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return errorResponse(401, fmt.Errorf("invalid token")), nil
		}

		ctx := context.WithValue(r.Context(), USER_EMAIL_KEY, claims["sub"].(string))
		ctx = context.WithValue(ctx, USER_NAME_KEY, claims["name"].(string))

		return next(fs, w, r.WithContext(ctx))
	}
}
