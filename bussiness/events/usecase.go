package events

import (
	"context"
	"time"
)

type eventsUsecase struct {
	eventsRepository 	Repository
	contextTimeout 		time.Duration
}

func NewEventsUsecase(timeout time.Duration, r Repository) Usecase {
	return &eventsUsecase{
		contextTimeout: 	timeout,
		eventsRepository: 	r,
	}
}

func (uc *eventsUsecase) Find(ctx context.Context, eventType string) ([]Domain, error) {
	resp, err := uc.eventsRepository.Find(ctx, eventType)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}