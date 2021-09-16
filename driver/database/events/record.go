package events

import (
	"auto_traveler/bussiness/events"
	"time"
)

type Events struct {
	ID 			int
	CreatedAt	time.Time
	UpdatedAt 	time.Time		
	Type		string		
	Name 		string	
	Description string	
	GoldReward 	int		
	XPReward	int	
}

// func fromDomain(domain *events.Domain) *Events {
// 	return &Events{
// 		Name:           domain.Name,
// 		Type: 			domain.Type,
// 		Description: 	domain.Description,
// 		GoldReward: 	domain.GoldReward,
// 		XPReward: 		domain.XPReward,
// 		CreatedAt:      domain.CreatedAt,
// 		UpdatedAt:      domain.UpdatedAt,
// 	}
// }

func (model *Events) ToDomain() (domain events.Domain) {
	if model != nil {
		domain = events.Domain{
			ID:				model.ID,
			CreatedAt:      model.CreatedAt,
			UpdatedAt:      model.UpdatedAt,
			Type: 			model.Type,
			Name:       	model.Name,
			Description: 	model.Description,
			GoldReward: 	domain.GoldReward,
			XPReward: 		domain.XPReward,	
		}
	}
	return domain
}
