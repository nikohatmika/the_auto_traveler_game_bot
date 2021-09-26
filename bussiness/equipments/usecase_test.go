package equipments_test

import (
	equipments "auto_traveler/bussiness/equipments"
	equipmentsMock "auto_traveler/bussiness/equipments/mocks"
	"auto_traveler/helper/messages"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	equipmentsRepository	equipmentsMock.Repository
	equipmentsUsecase 		equipments.Usecase
)

func setup() {
	equipmentsUsecase = equipments.NewEquipmentsUsecase(2, &equipmentsRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := equipments.Domain{
			ID:         1,
			Type:		"weapon",
			Name: 		"Name",
		}
		equipmentsRepository.On("FindByID", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := equipmentsUsecase.Find(context.Background(), "weapon")

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test case 2, invalid type", func(t *testing.T) {
		result, err := equipmentsUsecase.Find(context.Background(), "weapon")

		assert.Equal(t, result, equipments.Domain{})
		assert.Equal(t, err, messages.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		equipmentsRepository.On("FindByID", mock.AnythingOfType("int")).Return(equipments.Domain{}, errNotFound).Once()
		result, err := equipmentsUsecase.Find(context.Background(), "weapon")

		assert.Equal(t, result, equipments.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}
