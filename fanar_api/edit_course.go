package fanar

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Modalessi/iau_resources/models"
	"github.com/google/uuid"
)

type editCourseRequest struct {
	Title        string `json:"title,omitempty"`
	CourseCode   string `json:"course_code,omitempty"`
	Description  string `json:"description,omitempty"`
	CreditHours  *int   `json:"credit_hours,omitempty"`
	ContactHours *int   `json:"contact_hours,omitempty"`
}

type editCourseResponse struct {
	Jsonable     `json:"-"`
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Code         string    `json:"code"`
	Description  string    `json:"description"`
	CreditHours  int       `json:"credit_hours"`
	ContactHours int       `json:"contact_hours"`
}

func editCourse(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {

	courseData := editCourseRequest{}
	err := json.NewDecoder(r.Body).Decode(&courseData)
	if err != nil {
		return fanarMessageResponse(400, "invalid request, please check your json"), nil
	}
	defer r.Body.Close()

	courseID := r.URL.Query().Get("id")
	if courseID == "" {
		return fanarMessageResponse(400, "invalid request, please provide the course id as and id query"), nil
	}

	courseUUID, err := uuid.Parse(courseID)
	if err != nil {
		return fanarMessageResponse(400, "invalid course ID format"), nil
	}

	// Check if the course exists
	existingCourse, err := fs.Storage.GetCourseByID(r.Context(), courseID)
	if err != nil {
		return serverErrorResponse("db err: getting course from db"), err
	}

	if existingCourse == nil {
		return fanarMessageResponse(404, "course does not exist"), nil
	}

	// Start with existing course data and update only provided fields
	updatedCourse := &models.Course{
		ID:           &courseUUID,
		Title:        existingCourse.Title,
		Code:         existingCourse.Code,
		Description:  existingCourse.Description,
		CreditHours:  existingCourse.CreditHours,
		ContactHours: existingCourse.ContactHours,
	}

	// Update only the fields that were provided
	if courseData.Title != "" {
		updatedCourse.Title = courseData.Title
	}

	if courseData.CourseCode != "" {
		updatedCourse.Code = courseData.CourseCode
	}

	if courseData.Description != "" {
		updatedCourse.Description = courseData.Description
	}

	if courseData.CreditHours != nil {
		updatedCourse.CreditHours = *courseData.CreditHours
	}

	if courseData.ContactHours != nil {
		updatedCourse.ContactHours = *courseData.ContactHours
	}

	// Update the course in storage
	course, err := fs.Storage.UpdateCourse(r.Context(), updatedCourse)
	if err != nil {
		return serverErrorResponse(fmt.Sprintf("db err: updating course: %v", err)), err
	}

	res := &editCourseResponse{
		ID:           *course.ID,
		Title:        course.Title,
		Code:         course.Code,
		Description:  course.Description,
		CreditHours:  course.CreditHours,
		ContactHours: course.ContactHours,
	}

	return newFanarResponse(200, res), nil
}
