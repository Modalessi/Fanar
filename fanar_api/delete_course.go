package fanar

import (
	"net/http"

	"github.com/google/uuid"
)

type deleteCourseResponse struct {
	Jsonable `json:"-"`
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
}

func deleteCourse(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {

	courseID := r.URL.Query().Get("id")
	if courseID == "" {
		return fanarMessageResponse(401, "how the fuck do you wanna me know what course to delete"), nil
	}

	deletedCourse, err := fs.Storage.DeleteCourse(r.Context(), courseID)
	if err != nil {
		return nil, err
	}

	if deletedCourse == nil {
		return fanarMessageResponse(404, "course does not exist"), nil
	}

	res := &deleteCourseResponse{
		Id:    *deletedCourse.ID,
		Title: deletedCourse.Title,
	}

	return newFanarResponse(200, res), nil
}
