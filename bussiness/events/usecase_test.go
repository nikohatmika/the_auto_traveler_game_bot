package events_test

import (
	events "auto_traveler/bussiness/events"
	eventsMock "auto_traveler/bussiness/events/mocks"
	"auto_traveler/helper/messages"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	eventsRepository	eventsMock.Repository
	eventsUsecase 		events.Usecase
)

func setup() {
	eventsUsecase = events.NewEventsUsecase(2, &eventsRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := events.Domain{
			ID:         1,
			Type:		"expedition",
			Name: 		"Name",
		}
		eventsRepository.On("FindByID", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := eventsUsecase.Find(context.Background(), "expedition")

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test case 2, invalid type", func(t *testing.T) {
		result, err := eventsUsecase.Find(context.Background(), "expedition")

		assert.Equal(t, result, events.Domain{})
		assert.Equal(t, err, messages.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		eventsRepository.On("FindByID", mock.AnythingOfType("int")).Return(events.Domain{}, errNotFound).Once()
		result, err := eventsUsecase.Find(context.Background(), "expedition")

		assert.Equal(t, result, events.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}