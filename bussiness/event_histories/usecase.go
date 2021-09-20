package event_histories

import (
	"context"
	"time"
)

type eventHistoriesUsecase struct {
	eventHistoriesRepository 	Repository
	contextTimeout				time.Duration
}

func NewEventHistoriesUsecase(timeout time.Duration, r Repository) Usecase {
	return &eventHistoriesUsecase{
		contextTimeout: 			timeout,
		eventHistoriesRepository:	r,
	}
}

func (uc *eventHistoriesUsecase) Find(ctx context.Context, playerId int) ([]Domain, error) {
	resp, err := uc.eventHistoriesRepository.Find(ctx, playerId)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}