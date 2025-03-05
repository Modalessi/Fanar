package fanar

import "net/http"

func checkHealth(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {
	return fanarMessageResponse(200, "i am alive"), nil
}
