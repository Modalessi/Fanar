package fanar

import "net/http"

type FanarResponse struct {
	Jsonable
	Code     int      `json:"code"`
	Response Jsonable `json:"response"`
}

func newFanarResponse(code int, res Jsonable) *FanarResponse {
	return &FanarResponse{
		Code:     code,
		Response: res,
	}
}

func fanarMessageResponse(code int, msg string) *FanarResponse {
	res := struct {
		Jsonable `json:"-"`
		Code     int    `json:"code"`
		Message  string `json:"message"`
	}{
		Code:    code,
		Message: msg,
	}

	return &FanarResponse{
		Code:     code,
		Response: res,
	}
}

func errorResponse(code int, err error) *FanarResponse {

	res := struct {
		Jsonable `json:"-"`
		Code     int    `json:"code"`
		Message  string `json:"message"`
	}{
		Code:    code,
		Message: err.Error(),
	}

	return &FanarResponse{
		Code:     code,
		Response: res,
	}
}

func serverErrorResponse(msg string) *FanarResponse {
	res := struct {
		Jsonable `json:"-"`
		Code     int    `json:"code"`
		Message  string `json:"message"`
	}{
		Code:    500,
		Message: msg,
	}

	return &FanarResponse{
		Code:     500,
		Response: res,
	}
}

func respondWithJson(w http.ResponseWriter, code int, res Jsonable) {
	w.Header().Add("Content-Type", "application/json")

	data := &JsonWrapper{Data: res}

	w.WriteHeader(code)
	w.Write(data.JSON())
}
