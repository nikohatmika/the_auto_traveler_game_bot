package admin_test

import (
	admin "auto_traveler/bussiness/admin"
	adminMock "auto_traveler/bussiness/admin/mocks"
	"auto_traveler/helper/messages"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	adminRepository	adminMock.Repository
	adminUsecase 		admin.Usecase
)

func setup() {
	adminUsecase = admin.NewAdminUsecase(2, &adminRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := admin.Domain{
			ID:         1,
			Email:      "mail@gmail.com",
			Name: 		"Name",
		}
		adminRepository.On("FindByID", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := adminUsecase.FindByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.Email, result.Email)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := adminUsecase.FindByID(context.Background(), -1)

		assert.Equal(t, result, admin.Domain{})
		assert.Equal(t, err, messages.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		adminRepository.On("FindByID", mock.AnythingOfType("int")).Return(admin.Domain{}, errNotFound).Once()
		result, err := adminUsecase.FindByID(context.Background(), 10)

		assert.Equal(t, result, admin.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestFindByEmail(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := admin.Domain{
			ID:         1,
			Email:      "mail@gmail.com",
			Name: 		"Name",
		}
		adminRepository.On("FindByEmail", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := adminUsecase.FindByEmail(context.Background(), "mail@gmail.com")

		assert.Nil(t, err)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := adminUsecase.FindByEmail(context.Background(), "")

		assert.Equal(t, result, admin.Domain{})
		assert.Equal(t, err, messages.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) Email Not Found")
		adminRepository.On("FindByEmail", mock.AnythingOfType("int")).Return(admin.Domain{}, errNotFound).Once()
		result, err := adminUsecase.FindByEmail(context.Background(), "mail@gmail.com")

		assert.Equal(t, result, admin.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}