package fanar

import (
	"context"

	"github.com/Modalessi/iau_resources/models"
)

type Storage interface {
	DoesUserExistWithEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctc context.Context, id string) (*models.User, error)
	StoreUser(ctx context.Context, user *models.User) error
}
