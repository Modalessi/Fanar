package models

import (
	"net/url"

	"github.com/google/uuid"
)

// (Notes, Homeworks, Quizes, Labs, Slides, Midterms, Finals, Exams, OldExams)
var Tags = map[string]bool{
	"NOTES":     true,
	"HOMEWORKS": true,
	"QUIZZES":   true,
	"LABS":      true,
	"SLIDES":    true,
	"MIDTERMS":  true,
	"FINALS":    true,
	"EXAMS":     true,
	"OLDEXAMS":  true,
}

type Resource struct {
	ID          *uuid.UUID
	CourseID    uuid.UUID
	Title       string
	Description string
	FileExt     string
	Url         *url.URL
	Tags        []string
	Created_by  uuid.UUID
}

func NewResource(courseID uuid.UUID, title, description, fileExt string, tags []string, created_by uuid.UUID) *Resource {
	return &Resource{
		CourseID:    courseID,
		Title:       title,
		Description: description,
		FileExt:     fileExt,
		Tags:        tags,
		Created_by:  created_by,
	}
}

func InvalidResourceTags(tags ...string) []string {
	t := []string{}

	for _, v := range tags {
		if valid := Tags[v]; !valid {
			t = append(t, v)
		}
	}

	return t
}
