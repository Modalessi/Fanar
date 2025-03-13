package fanar

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func NewAdminOnlyMiddlware(adminEmail string) func(next FanaerHandler) FanaerHandler {
	return func(next FanaerHandler) FanaerHandler {
		return func(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {
			email, ok := r.Context().Value(USER_EMAIL_KEY).(string)
			if !ok {
				return serverErrorResponse("server error: invalid context"), fmt.Errorf("invalid email type in context")
			}

			if email == "" {
				return errorResponse(401, fmt.Errorf("unauthorized: authentication required")), nil
			}

			if !strings.EqualFold(email, adminEmail) {
				log.Printf("the user email is: %v\nthe admin email is: %v", email, adminEmail)
				return errorResponse(403, fmt.Errorf("forbidden: admin access required")), nil
			}

			return next(fs, w, r)
		}
	}
}
