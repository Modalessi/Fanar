package fanar

import (
	"net/http"

	"github.com/google/uuid"
)

type getCourseResourcesResponse struct {
	Jsonable      `json:"-"`
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Tags          []string  `json:"tags"`
	ByCurrentUser bool      `json:"by_current_user"`
}

func getCourseResources(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {
	courseID := r.URL.Query().Get("id")
	if courseID == "" {
		return fanarMessageResponse(400, "course id was not provided"), nil
	}

	course, err := fs.Storage.GetCourseByID(r.Context(), courseID)
	if err != nil {
		return serverErrorResponse("error getting course from db"), err
	}

	if course == nil {
		return fanarMessageResponse(404, "course does not exist"), nil
	}

	resources, err := fs.Storage.GetCourseResources(r.Context(), courseID)

	res := make([]getCourseResourcesResponse, len(resources))

	for i, v := range resources {
		// TODO: By Current user, works
		res[i] = getCourseResourcesResponse{
			ID:            *v.ID,
			Title:         v.Title,
			Description:   v.Description,
			Tags:          v.Tags,
			ByCurrentUser: false,
		}
	}

	return newFanarResponse(200, &JsonWrapper{Data: res}), nil
}
