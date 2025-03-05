package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Modalessi/iau_resources/database"
	"github.com/Modalessi/iau_resources/models"
	"github.com/google/uuid"
)

func (s *Storage) StoreCourse(ctx context.Context, course *models.Course) error {
	createCourseParams := database.CreateCourseParams{
		Title:        course.Title,
		CourseCode:   course.Code,
		Description:  course.Description,
		CreditHours:  int32(course.CreditHours),
		ContactHours: int32(course.ContactHours),
	}

	dbCourse, err := s.queries.CreateCourse(ctx, createCourseParams)
	if err != nil {
		return fmt.Errorf("db error: storing couse: %v", err)
	}

	course.ID = &dbCourse.ID
	return nil
}

func (s *Storage) GetCourseByID(ctx context.Context, id string) (*models.Course, error) {

	courseUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid course id: %v", err)
	}

	dbCourse, err := s.queries.GetCourseByID(ctx, courseUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("db error: getting course from db: %v", err)
	}

	return &models.Course{
		ID:           &dbCourse.ID,
		Title:        dbCourse.Title,
		Description:  dbCourse.Description,
		Code:         dbCourse.CourseCode,
		CreditHours:  int(dbCourse.CreditHours),
		ContactHours: int(dbCourse.ContactHours),
	}, nil
}
