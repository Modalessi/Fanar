package fanar

import (
	"encoding/json"
	"net/http"

	"github.com/Modalessi/iau_resources/models"
	"github.com/google/uuid"
)

type createCourseRequest struct {
	Title        string `json:"title" validate:"required"`
	CourseCode   string `json:"course_code" validate:"required"`
	Description  string `json:"description" validate:"required"`
	CreditHours  int    `json:"credit_hours" validate:"required"`
	ContactHours int    `json:"contact_hours" validate:"required"`
}

type createCourseResponse struct {
	Jsonable `json:"-"`
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
}

func createCourse(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {

	courseData := createCourseRequest{}
	err := json.NewDecoder(r.Body).Decode(&courseData)
	if err != nil {
		return fanarMessageResponse(401, "invalid request, please check your json"), nil
	}
	defer r.Body.Close()

	course := models.NewCourse(courseData.Title, courseData.CourseCode, courseData.CreditHours, courseData.ContactHours)
	course.SetDescription(courseData.Description)

	err = fs.Storage.StoreCourse(r.Context(), course)
	if err != nil {
		return fanarMessageResponse(401, "invalid requst, this course already exist"), nil
	}

	res := &createCourseResponse{
		Id:    *course.ID,
		Title: course.Title,
	}

	return newFanarResponse(201, res), nil
}
