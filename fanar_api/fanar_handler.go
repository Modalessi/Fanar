package fanar

import (
	"log"
	"net/http"
	"time"
)

type FanaerHandler func(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error)

func withServer(fs *FanarServer, handler FanaerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		res, err := handler(fs, w, r)
		duration := time.Since(start)
		if err != nil {
			log.Printf("Error handling %s %s: %v", r.Method, r.URL.Path, err)
		}

		if res != nil {
			respondWithJson(w, res.Code, res.Response)
		} else {
			serverErrRes := serverErrorResponse("something went wrong in the server")
			respondWithJson(w, 500, serverErrRes)
		}

		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
	}
}
