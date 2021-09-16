package request

import (
	"auto_traveler/bussiness/events"
	"time"
)

type Events struct {
	ID 			int 	`json:"id"`
	CreatedAt	string 	`json:"created_at"`
	UpdatedAt 	string	`json:"updated_at"`	
	Type		string	`json:"type"`
	Name 		string	`json:"name"`
	Description string	`json:"description"`
	GoldReward 	int		`json:"gold_reward"`	
	XPReward	int		`json:"xp_reward"`
}

func (req *Events) ToDomain() (res *events.Domain, err error) {
	return &events.Domain{
		ID:				req.ID,		 		
		CreatedAt:  	time.Now().UTC(),
		UpdatedAt:  	time.Now().UTC(),
		Type: 			req.Type,
		Name:			req.Name,
		Description: 	req.Description, 
		GoldReward: 	req.GoldReward,
		XPReward: 		req.XPReward,
	}, err
}