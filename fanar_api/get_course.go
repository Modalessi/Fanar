package fanar

import (
	"net/http"

	"github.com/google/uuid"
)

type getCourseResponse struct {
	Jsonable     `json:"-"`
	ID           uuid.UUID
	Title        string
	Code         string
	Description  string
	CreditHours  int
	ContactHours int
}

func getCourse(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {
	courseID := r.URL.Query().Get("id")
	if courseID == "" {
		return fanarMessageResponse(400, "invalid request, please provide id query for the course id"), nil
	}

	course, err := fs.Storage.GetCourseByID(r.Context(), courseID)
	if err != nil {
		return serverErrorResponse("db err: getting course from db"), err
	}

	if course == nil {
		return fanarMessageResponse(404, "course does not exist"), nil
	}

	res := &getCourseResponse{
		ID:           *course.ID,
		Title:        course.Title,
		Code:         course.Code,
		Description:  course.Description,
		CreditHours:  course.CreditHours,
		ContactHours: course.ContactHours,
	}

	return newFanarResponse(200, res), nil
}
