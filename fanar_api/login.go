package fanar

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Jsonable `json:"_"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func login(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {

	loginData := loginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		return fanarMessageResponse(400, "invalid request, please check you sent a valid json"), nil
	}
	defer r.Body.Close()

	user, err := fs.Storage.GetUserByEmail(r.Context(), loginData.Email)
	if err != nil {
		return serverErrorResponse("db err: getting user from databae"), err
	}

	if user == nil {
		return fanarMessageResponse(404, "user does not exist"), nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return fanarMessageResponse(401, "wrong credintals, please check your password"), nil
	}

	token, err := NewJWTTokenWithClaims(user.Name, user.Email, fs.JWTSecret)
	if err != nil {
		return serverErrorResponse("something went wrong"), err
	}

	res := loginResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return newFanarResponse(200, res), nil

}
