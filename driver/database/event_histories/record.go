package event_histories

import (
	"auto_traveler/bussiness/event_histories"
	"time"
)

type EventHistories struct {
	ID 			int
	CreatedAt	time.Time
	PlayerID	int
	EventID		int
	Result		bool	
}

func (model *EventHistories) ToDomain() (domain event_histories.Domain) {
	if model != nil {
		domain = event_histories.Domain{
			ID:			model.ID,
			CreatedAt:	model.CreatedAt,
			PlayerID:	model.PlayerID,
			EventID: 	model.EventID,
			Result:     model.Result,
		}
	}
	return domain
}
