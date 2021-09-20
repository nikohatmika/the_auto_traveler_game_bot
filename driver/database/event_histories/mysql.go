package event_histories

import (
	"auto_traveler/bussiness/event_histories"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type eventHistoriesRepository struct {
	conn *gorm.DB
}

func NewEventHistoriesRepository(conn *gorm.DB) event_histories.Repository {
	return &eventHistoriesRepository{
		conn: conn,
	}
}

func (r *eventHistoriesRepository) Find(ctx context.Context, playerId int) ([]event_histories.Domain, error) {
	res := []EventHistories{}
	fmt.Println("---- RRR ...", r)


	query := r.conn

	if playerId != 0 {
		query = query.Where("player_id", playerId)
	}

	fmt.Println("Q U E R Y...", query)

	err := query.Find(&res).Error
	if err != nil {
		return []event_histories.Domain{}, err
	}

	eventHistoriesDomain := []event_histories.Domain{}
	for _, value := range res {
		eventHistoriesDomain = append(eventHistoriesDomain, value.ToDomain())
	}

	fmt.Println("res", res)
	// fmt.Println("query", query)

	return eventHistoriesDomain, nil
}