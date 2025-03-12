package fanar

import (
	"fmt"
	"net/http"
)

func getResourceLink(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {
	resourceID := r.URL.Query().Get("id")
	if resourceID == "" {
		return fanarMessageResponse(400, "please pass the resource id as query"), nil
	}

	resource, err := fs.Storage.GetResource(r.Context(), resourceID)
	if err != nil {
		return serverErrorResponse(fmt.Sprintf("db error, getting resoure with id %s: %s", resourceID, err)), err
	}

	downladURL, err := fs.Storage.GetResourceDownloadURL(r.Context(), resource)
	if err != nil {
		return serverErrorResponse(fmt.Sprintf("error generating downladable link %v", err)), err
	}

	return fanarMessageResponse(200, downladURL), nil
}
