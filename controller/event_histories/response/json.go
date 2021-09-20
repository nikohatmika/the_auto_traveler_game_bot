package response

import "auto_traveler/bussiness/event_histories"

type EventHistories struct {
	ID 			int 	`json:"id"`
	CreatedAt	string 	`json:"created_at"`
	PlayerID 	int		`json:"player_id"`	
	EventID		int		`json:"event_id"`
	Result 		bool	`json:"result"`
}

func FromDomain(domain *event_histories.Domain) (res *EventHistories) {
	if domain != nil {
		res = &EventHistories{
			ID:      	domain.ID,
			CreatedAt:  domain.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
			PlayerID: 	domain.PlayerID,
			EventID:	domain.EventID,
			Result: 	domain.Result, 
		}
	}

	return res
}