package fanar

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/Modalessi/iau_resources/models"
	"github.com/google/uuid"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	Jsonable `json:"-"`
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
}

func register(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {

	registerData := registerRequest{}
	err := json.NewDecoder(r.Body).Decode(&registerData)
	defer r.Body.Close()

	if err != nil {
		return fanarMessageResponse(400, "invalid request, please check you sent valid json"), nil
	}

	if !isValidCredentials(registerData.Email, registerData.Password) {
		return fanarMessageResponse(400, "invalid request, please check you entered valid credintals"), nil
	}

	emailTaken, err := fs.Storage.DoesUserExistWithEmail(r.Context(), registerData.Email)
	if err != nil {
		return serverErrorResponse(err.Error()), err
	}

	if emailTaken {
		return fanarMessageResponse(400, "invalid request, email is already taken"), nil
	}

	newUser := models.NewUser(registerData.Name, registerData.Email, registerData.Password)
	err = fs.Storage.StoreUser(r.Context(), newUser)
	if err != nil {
		return serverErrorResponse("error creating new user"), err
	}

	res := registerResponse{
		Id:    *newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return newFanarResponse(201, res), nil
}

func isValidCredentials(email string, pw string) bool {

	emailRegex := `^[a-zA-Z0-9._%+-]+@iau.edu.sa$`
	isValidEmail, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false
	}
	if !isValidEmail {
		return false
	}

	if len([]byte(pw)) > 72 {
		return false
	}

	if len(pw) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pw)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(pw)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(pw)

	return hasUpper && hasLower && hasNumber
}
