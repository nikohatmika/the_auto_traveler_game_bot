package events

import (
	"auto_traveler/bussiness/events"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type eventsRepository struct {
	conn *gorm.DB
}

func NewEventsRepository(conn *gorm.DB) events.Repository {
	return &eventsRepository{
		conn: conn,
	}
}

func (r *eventsRepository) Find(ctx context.Context, eventType string) ([]events.Domain, error) {
	res := []Events{}

	query := r.conn

	if eventType != "" {
		query = query.Where("type", eventType)
	}

	err := query.Find(&res).Error
	if err != nil {
		return []events.Domain{}, err
	}

	eventsDomain := []events.Domain{}
	for _, value := range res {
		eventsDomain = append(eventsDomain, value.ToDomain())
	}

	fmt.Println("res", res)
	// fmt.Println("query", query)

	return eventsDomain, nil
}