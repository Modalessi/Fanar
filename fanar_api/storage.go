package fanar

import (
	"context"
	"io"

	"github.com/Modalessi/iau_resources/models"
)

type Storage interface {
	// users
	DoesUserExistWithEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctc context.Context, email string) (*models.User, error)
	StoreUser(ctx context.Context, user *models.User) error

	// courses
	GetCourseByID(ctx context.Context, id string) (*models.Course, error)
	StoreCourse(ctx context.Context, course *models.Course) error
	DeleteCourse(ctx context.Context, id string) (*models.Course, error)
	UpdateCourse(ctx context.Context, course *models.Course) (*models.Course, error)

	// resources
	StoreResource(ctx context.Context, resource *models.Resource, file io.Reader, contentType string) error
	GetResource(ctx context.Context, id string) (*models.Resource, error)
	GetCourseResources(ctx context.Context, courseID string) ([]*models.Resource, error)
	GetResourceDownloadURL(ctx context.Context, resource *models.Resource) (string, error)
}
