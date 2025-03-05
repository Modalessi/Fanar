package fanar

import (
	"net/http"

	"github.com/Modalessi/iau_resources/utils"
)

type protectedResponse struct {
	Jsonable `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func protected(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {

	userName, ok := r.Context().Value(USER_NAME_KEY).(string)
	utils.Assert(ok, "the value of the USER_NAME_KEY was not string, there is something deeply wrong here")

	userEmail, ok := r.Context().Value(USER_EMAIL_KEY).(string)
	utils.Assert(ok, "the value of the USER_EMAIL_KEY was not string, there is something deeply wrong here")

	res := protectedResponse{
		Name:  userName,
		Email: userEmail,
	}

	return newFanarResponse(200, res), nil

}
