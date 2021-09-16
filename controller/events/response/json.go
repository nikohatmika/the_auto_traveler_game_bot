package response

import "auto_traveler/bussiness/events"

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

func FromDomain(domain *events.Domain) (res *Events) {
	if domain != nil {
		res = &Events{
			ID:      		domain.ID,
			CreatedAt:      domain.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
			UpdatedAt:      domain.UpdatedAt.UTC().Format("2006-01-02 15:04:05"),
			Type: 			domain.Type,
			Name:         	domain.Name,
			Description:  	domain.Description,
			GoldReward:     domain.GoldReward,
			XPReward: 		domain.XPReward,
		}
	}

	return res
}