package services

import (
	"UserManagement/commons/apploggers"
	"UserManagement/internals/db"
	dbModels "UserManagement/internals/db/models"
	"UserManagement/internals/models"
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Services", func() {

	Describe("GetUserById", func() {
		It("empty userid", func() {
			var resUser *models.User = nil
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeGetUserById: func(ctx context.Context, id string) (*models.User, error) {
					return resUser, fmt.Errorf("invalid/empty userid")
				},
			}
			eservice := NewUserEventService(dbservice)
			user, err := eservice.GetUserById(ctx, "")
			Expect(err).To(HaveOccurred())
			Expect(user).To(Equal(resUser))
			Expect(err.Error()).To(Equal("invalid/empty userid"))
		})

		It("valid", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeGetUserById: func(context context.Context, userId string) (*models.User, error) {
					return &models.User{
						Name:     "John",
						Email:    "testuser@test.com",
						Age:      27,
						IsActive: true,
						Type:     "Customer",
					}, nil
				},
			}

			eservice := NewUserEventService(dbservice)
			user, err := eservice.GetUserById(ctx, "123")

			Expect(err).NotTo(HaveOccurred())
			Expect(user).NotTo(BeNil())
			Expect(user.Name).To(Equal("John"))
			Expect(user.Email).To(Equal("testuser@test.com"))
		})
	})

	Describe("DeleteUserById", func() {
		It("empty userid", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeDeleteUserById: func(ctx context.Context, id string) error {
					return fmt.Errorf("invalid/empty userid")
				},
			}
			eservice := NewUserEventService(dbservice)
			err := eservice.DeleteUserById(ctx, "")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("invalid/empty userid"))
		})

		It("valid", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeDeleteUserById: func(context context.Context, userId string) error {
					return nil
				},
			}

			eservice := NewUserEventService(dbservice)
			err := eservice.DeleteUserById(ctx, "123")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("GetUsers", func() {
		var resUser []*models.User = nil
		It("dberror", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeGetUsers: func(ctx context.Context) ([]*models.User, error) {
					return resUser, fmt.Errorf("dberror")
				},
			}
			eservice := NewUserEventService(dbservice)
			users, err := eservice.GetUsers(ctx)
			Expect(err).To(HaveOccurred())
			Expect(users).To(Equal(resUser))
			Expect(err.Error()).To(Equal("dberror"))
		})

		It("valid", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeGetUsers: func(ctx context.Context) ([]*models.User, error) {
					return []*models.User{
						{
							Name:     "John",
							Email:    "testuser@test.com",
							Age:      27,
							IsActive: true,
							Type:     "Customer",
						},
					}, nil
				},
			}

			eservice := NewUserEventService(dbservice)
			users, err := eservice.GetUsers(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(users).NotTo(BeNil())
			Expect(len(users)).To(Equal(1))
			Expect(users[0].Name).To(Equal("John"))
			Expect(users[0].Email).To(Equal("testuser@test.com"))
		})
	})

	Describe("UpdateUser", func() {
		It("dberror", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeUpdateUser: func(ctx context.Context, user *dbModels.UserSchema, userId string) error {
					return fmt.Errorf("dberror")
				},
			}

			eservice := NewUserEventService(dbservice)
			err := eservice.UpdateUser(ctx, &models.User{
				Name:     "John",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			}, "123")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("dberror"))
		})

		It("valid", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeUpdateUser: func(ctx context.Context, user *dbModels.UserSchema, userId string) error {
					return nil
				},
			}

			eservice := NewUserEventService(dbservice)
			err := eservice.UpdateUser(ctx, &models.User{
				Name:     "John",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			}, "123")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("CreateUser", func() {
		It("dberror", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeSaveUser: func(ctx context.Context, user *dbModels.UserSchema) (string, error) {
					return "", fmt.Errorf("dberror")
				},
			}

			eservice := NewUserEventService(dbservice)
			users, err := eservice.CreateUser(ctx, &models.User{
				Name:     "John",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("dberror"))
			Expect(len(users)).To(Equal(0))
		})

		It("valid", func() {
			ctx, _ := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
			dbservice := db.MockDBService{
				FakeSaveUser: func(ctx context.Context, user *dbModels.UserSchema) (string, error) {
					return "123", nil
				},
			}

			eservice := NewUserEventService(dbservice)
			users, err := eservice.CreateUser(ctx, &models.User{
				Name:     "John",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(len(users)).NotTo(Equal(0))
		})
	})

})
