package apis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"UserManagement/commons"
	"UserManagement/internals/models"
	"UserManagement/internals/services"
)

var _ = Describe("User API Controller", func() {

	Describe("GetUserById", func() {
		It("empty userid", func() {
			eservice := services.MockEventService{}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("")

			controller := NewUserController(eservice)
			err := controller.GetUserById(c)
			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("'id' is required"))
		})

		It("invalid userid", func() {
			eservice := services.MockEventService{
				FakeGetUserById: func(context context.Context, userId string) (*models.User, error) {
					return nil, fmt.Errorf("invalid 'id' provided")
				},
			}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.GetUserById(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("invalid 'id' provided"))
		})

		It("valid", func() {
			eservice := services.MockEventService{
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
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.GetUserById(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var response *models.User
			uerror := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerror).NotTo(HaveOccurred())
			Expect(response.Name).To(Equal("John"))
			Expect(response.Email).To(Equal("testuser@test.com"))
		})
	})

	Describe("DeleteUserById", func() {
		It("empty userid", func() {
			eservice := services.MockEventService{}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/users/123", nil)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("")

			controller := NewUserController(eservice)
			err := controller.DeleteUserById(c)
			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("'id' is required"))
		})

		It("invalid userid", func() {
			eservice := services.MockEventService{
				FakeDeleteUserById: func(context context.Context, userId string) error {
					return fmt.Errorf("invalid 'id' provided")
				},
			}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/users/123", nil)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.DeleteUserById(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("invalid 'id' provided"))
		})

		It("valid", func() {
			eservice := services.MockEventService{
				FakeDeleteUserById: func(context context.Context, userId string) error {
					return nil
				},
			}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/users/123", nil)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.DeleteUserById(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusNoContent))
		})
	})

	Describe("GetUsers", func() {
		It("invalid userid", func() {
			eservice := services.MockEventService{
				FakeGetUsers: func(context context.Context) ([]*models.User, error) {
					return nil, fmt.Errorf("invalid 'id' provided")
				},
			}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			c := e.NewContext(req, rec)

			controller := NewUserController(eservice)
			err := controller.GetUsers(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("invalid 'id' provided"))
		})

		It("valid", func() {
			id, _ := primitive.ObjectIDFromHex("66cf79ff71d915dbe917ea3a")
			eservice := services.MockEventService{
				FakeGetUsers: func(context context.Context) ([]*models.User, error) {
					return []*models.User{
						{
							Id:       id,
							Name:     "John",
							Email:    "testuser@test.com",
							Age:      27,
							IsActive: true,
							Type:     "Customer",
						},
					}, nil
				},
			}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			c := e.NewContext(req, rec)

			controller := NewUserController(eservice)
			err := controller.GetUsers(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var response map[string]interface{}
			fmt.Print(rec.Body.String())
			uerror := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerror).NotTo(HaveOccurred())
			Expect(response["total"].(float64)).To(Equal(float64(1)))
			Expect(response["users"]).NotTo(BeNil())

			pbytes, _ := json.Marshal(response["users"])
			var users []*models.User
			uerror = json.Unmarshal(pbytes, &users)
			Expect(uerror).NotTo(HaveOccurred())
			Expect(len(users)).To(Equal(1))
			Expect(users[0].Name).To(Equal("John"))
			Expect(users[0].Email).To(Equal("testuser@test.com"))
		})
	})

	Describe("CreateUser", func() {
		It("empty payload", func() {
			eservice := services.MockEventService{
				FakeCreateUser: func(context context.Context, user *models.User) (string, error) {
					return "", fmt.Errorf("invalid 'id' provided")
				},
			}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", nil)
			c := e.NewContext(req, rec)

			controller := NewUserController(eservice)
			err := controller.CreateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("invalid request payload"))
		})

		It("invalid payload", func() {
			eservice := services.MockEventService{
				FakeCreateUser: func(context context.Context, user *models.User) (string, error) {
					return "", fmt.Errorf("invalid 'id' provided")
				},
			}

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader("test"))
			c := e.NewContext(req, rec)

			controller := NewUserController(eservice)
			err := controller.CreateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("invalid request payload"))
		})

		It("empty name", func() {
			eservice := services.MockEventService{
				FakeCreateUser: func(context context.Context, user *models.User) (string, error) {
					return "test-id", nil
				},
			}
			pbytes, merr := json.Marshal(&models.User{
				Name:     "",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			})
			Expect(merr).To(BeNil())
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(pbytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.CreateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("'name' is required"))
		})

		It("empty email", func() {
			eservice := services.MockEventService{
				FakeCreateUser: func(context context.Context, user *models.User) (string, error) {
					return "test-id", nil
				},
			}
			pbytes, merr := json.Marshal(&models.User{
				Name:     "John",
				Email:    "",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			})
			Expect(merr).To(BeNil())
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(pbytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.CreateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("'email' is required"))
		})

		It("dberror", func() {
			eservice := services.MockEventService{
				FakeCreateUser: func(context context.Context, user *models.User) (string, error) {
					return "test-id", fmt.Errorf("dberror")
				},
			}
			pbytes, merr := json.Marshal(&models.User{
				Name:     "John",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			})
			Expect(merr).To(BeNil())
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(pbytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.CreateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("dberror"))
		})

		It("valid", func() {
			eservice := services.MockEventService{
				FakeCreateUser: func(context context.Context, user *models.User) (string, error) {
					return "test-id", nil
				},
			}
			pbytes, merr := json.Marshal(&models.User{
				Name:     "John",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			})
			Expect(merr).To(BeNil())
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(pbytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.CreateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusCreated))

			var response map[string]interface{}
			uerror := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerror).NotTo(HaveOccurred())
			Expect(response["id"]).To(Equal("test-id"))
		})

	})

	Describe("UpdateUser", func() {
		It("empty userid", func() {
			eservice := services.MockEventService{}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("")

			controller := NewUserController(eservice)
			err := controller.UpdateUser(c)
			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("'id' is required"))
		})

		It("empty payload", func() {
			eservice := services.MockEventService{
				FakeUpdateUser: func(context context.Context, user *models.User, userId string) error {
					return nil
				},
			}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.UpdateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("invalid request payload"))
		})

		It("invalid payload", func() {
			eservice := services.MockEventService{
				FakeUpdateUser: func(context context.Context, user *models.User, userId string) error {
					return nil
				},
			}
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users/1", strings.NewReader("test"))
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.UpdateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("invalid request payload"))
		})

		It("empty name", func() {
			eservice := services.MockEventService{
				FakeUpdateUser: func(context context.Context, user *models.User, userId string) error {
					return nil
				},
			}
			user := &models.User{
				Name:     "",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			}
			pbytes, merr := json.Marshal(user)
			Expect(merr).To(BeNil())

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(pbytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.UpdateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("'name' is required"))
		})

		It("empty email", func() {
			eservice := services.MockEventService{
				FakeUpdateUser: func(context context.Context, user *models.User, userId string) error {
					return nil
				},
			}
			user := &models.User{
				Name:     "John",
				Email:    "",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			}
			pbytes, merr := json.Marshal(user)
			Expect(merr).To(BeNil())

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(pbytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.UpdateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("'email' is required"))
		})

		It("invalid userid", func() {
			eservice := services.MockEventService{
				FakeUpdateUser: func(context context.Context, user *models.User, userId string) error {
					return fmt.Errorf("invalid 'id' provided")
				},
			}
			user := &models.User{
				Name:     "John",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			}
			pbytes, merr := json.Marshal(user)
			Expect(merr).To(BeNil())

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(pbytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.UpdateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var response *commons.ApiErrorResponsePayload
			uerr := json.Unmarshal(rec.Body.Bytes(), &response)
			Expect(uerr).NotTo(HaveOccurred())
			Expect(response.Status).To(Equal("Error"))
			Expect(response.Message).To(Equal("invalid 'id' provided"))
		})

		It("valid", func() {
			eservice := services.MockEventService{
				FakeUpdateUser: func(context context.Context, user *models.User, userId string) error {
					return nil
				},
			}
			user := &models.User{
				Name:     "John",
				Email:    "testuser@test.com",
				Age:      27,
				IsActive: true,
				Type:     "Customer",
			}
			pbytes, merr := json.Marshal(user)
			Expect(merr).To(BeNil())

			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(pbytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("123")

			controller := NewUserController(eservice)
			err := controller.UpdateUser(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))
		})
	})

})
