package request

import (
	"auto_traveler/bussiness/event_histories"
	"time"
)

type EventHistories struct {
	ID 			int 	`json:"id"`
	CreatedAt	string 	`json:"created_at"`
	PlayerID 	int	`json:"player_id"`	
	EventID		int	`json:"event_id"`
	Result 		bool	`json:"result"`
}

func (req *EventHistories) ToDomain() (res *event_histories.Domain, err error) {
	return &event_histories.Domain{
		ID:			req.ID,		 		
		CreatedAt:  time.Now().UTC(),
		PlayerID: 	req.PlayerID,
		EventID:	req.EventID,
		Result: 	req.Result, 
	}, err
}