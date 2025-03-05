package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Modalessi/iau_resources/database"
	"github.com/Modalessi/iau_resources/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) StoreUser(ctx context.Context, user *models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("db error: genrating hashed password returned an error %v", err)
	}

	createUserParams := database.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	}
	dbUser, err := s.queries.CreateUser(ctx, createUserParams)
	if err != nil {
		return fmt.Errorf("db error: creating user in database %v", err)
	}

	user.ID = &dbUser.ID
	user.Password = dbUser.Password
	return nil
}

func (s *Storage) DoesUserExistWithEmail(ctx context.Context, email string) (bool, error) {

	exist, err := s.queries.DoesUserExistByEmail(ctx, email)
	if err != nil {
		return false, fmt.Errorf("db error: error checking if a user exist with email %v", err)
	}

	return exist, nil
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	userDB, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("db error: %v", err)
	}

	return &models.User{
		ID:       &userDB.ID,
		Name:     userDB.Name,
		Email:    userDB.Email,
		Password: userDB.Password,
	}, nil
}
