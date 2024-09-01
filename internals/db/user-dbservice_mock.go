package db

import (
	dbmodel "UserManagement/internals/db/models"
	"UserManagement/internals/models"
	"context"
	"fmt"
)

type MockDBService struct {
	FakeGetUserById    func(ctx context.Context, id string) (*models.User, error)
	FakeDeleteUserById func(ctx context.Context, id string) error
	FakeGetUsers       func(ctx context.Context) ([]*models.User, error)
	FakeSaveUser       func(ctx context.Context, user *dbmodel.UserSchema) (string, error)
	FakeUpdateUser     func(ctx context.Context, user *dbmodel.UserSchema, userId string) error
}

func (m MockDBService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	if m.FakeGetUserById != nil {
		return m.FakeGetUserById(ctx, id)
	}
	return nil, fmt.Errorf("GetUserById-error")
}

func (m MockDBService) DeleteUserById(ctx context.Context, id string) error {
	if m.FakeDeleteUserById != nil {
		return m.FakeDeleteUserById(ctx, id)
	}
	return fmt.Errorf("DeleteUserById-error")
}

func (m MockDBService) GetUsers(ctx context.Context) ([]*models.User, error) {
	if m.FakeGetUsers != nil {
		return m.FakeGetUsers(ctx)
	}
	return nil, fmt.Errorf("GetUsers-error")
}

func (m MockDBService) SaveUser(ctx context.Context, user *dbmodel.UserSchema) (string, error) {
	if m.FakeSaveUser != nil {
		return m.FakeSaveUser(ctx, user)
	}
	return "", fmt.Errorf("SaveUser-error")
}

func (m MockDBService) UpdateUser(ctx context.Context, user *dbmodel.UserSchema, userId string) error {
	if m.FakeUpdateUser != nil {
		return m.FakeUpdateUser(ctx, user, userId)
	}
	return fmt.Errorf("UpdateUser-error")
}
