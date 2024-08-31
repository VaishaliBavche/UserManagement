package services

import (
	"UserManagement/internals/models"
	"context"
	"fmt"
)

type MockEventService struct {
	FakeGetUserById    func(context context.Context, userId string) (*models.User, error)
	FakeDeleteUserById func(context context.Context, userId string) error
	FakeGetUsers       func(context context.Context) ([]*models.User, error)
	FakeCreateUser     func(context context.Context, user *models.User) (string, error)
	FakeUpdateUser     func(context context.Context, user *models.User, userId string) error
}

func (m MockEventService) CreateUser(context context.Context, user *models.User) (string, error) {
	if m.FakeCreateUser != nil {
		return m.FakeCreateUser(context, user)
	}
	return "", fmt.Errorf("CreateUser-error")
}

func (m MockEventService) UpdateUser(context context.Context, user *models.User, userId string) error {
	if m.FakeUpdateUser != nil {
		return m.FakeUpdateUser(context, user, userId)
	}
	return fmt.Errorf("UpdateUser-error")
}

func (m MockEventService) GetUsers(context context.Context) ([]*models.User, error) {
	if m.FakeGetUsers != nil {
		return m.FakeGetUsers(context)
	}
	return nil, fmt.Errorf("GetUsers-error")
}

func (m MockEventService) DeleteUserById(context context.Context, userId string) error {
	if m.FakeDeleteUserById != nil {
		return m.FakeDeleteUserById(context, userId)
	}
	return fmt.Errorf("DeleteUserById-error")
}

func (m MockEventService) GetUserById(context context.Context, userId string) (*models.User, error) {
	if m.FakeGetUserById != nil {
		return m.FakeGetUserById(context, userId)
	}
	return nil, fmt.Errorf("GetUserById-error")
}
