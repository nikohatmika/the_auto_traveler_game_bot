package players_test

import (
	players "auto_traveler/bussiness/players"
	playersMock "auto_traveler/bussiness/players/mocks"
	"auto_traveler/helper/messages"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	playersRepository	playersMock.Repository
	playersUsecase 		players.Usecase
)

func setup() {
	playersUsecase = players.NewPlayerUsecase(2, &playersRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := players.Domain{
			ID:         1,
			Email:      "mail@gmail.com",
			Name: 		"Name",
			Level:      1,
		}
		playersRepository.On("FindByID", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := playersUsecase.FindByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.Email, result.Email)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := playersUsecase.FindByID(context.Background(), -1)

		assert.Equal(t, result, players.Domain{})
		assert.Equal(t, err, messages.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		playersRepository.On("FindByID", mock.AnythingOfType("int")).Return(players.Domain{}, errNotFound).Once()
		result, err := playersUsecase.FindByID(context.Background(), 10)

		assert.Equal(t, result, players.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestFindByEmail(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := players.Domain{
			ID:         1,
			Email:      "mail@gmail.com",
			Name: 		"Name",
			Level:      1,
		}
		playersRepository.On("FindByEmail", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := playersUsecase.FindByEmail(context.Background(), "mail@gmail.com")

		assert.Nil(t, err)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := playersUsecase.FindByEmail(context.Background(), "")

		assert.Equal(t, result, players.Domain{})
		assert.Equal(t, err, messages.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) Email Not Found")
		playersRepository.On("FindByEmail", mock.AnythingOfType("int")).Return(players.Domain{}, errNotFound).Once()
		result, err := playersUsecase.FindByEmail(context.Background(), "mail@gmail.com")

		assert.Equal(t, result, players.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}