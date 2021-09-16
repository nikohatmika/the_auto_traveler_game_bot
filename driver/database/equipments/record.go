package equipments

import (
	"auto_traveler/bussiness/equipments"
	"time"
)

type Equipments struct {
	ID 			int
	CreatedAt	time.Time
	UpdatedAt 	time.Time		
	Type		string		
	Name 		string	
	Description string	
	ATK 		int			
	DEF			int		
	HP			int		
}

// func fromDomain(domain *equipments.Domain) *Equipments {
// 	return &Equipments{
// 		Name:           domain.Name,
// 		Type: 			domain.Type,
// 		Description: 	domain.Description,
// 		GoldReward: 	domain.GoldReward,
// 		XPReward: 		domain.XPReward,
// 		CreatedAt:      domain.CreatedAt,
// 		UpdatedAt:      domain.UpdatedAt,
// 	}
// }

func (model *Equipments) ToDomain() (domain equipments.Domain) {
	if model != nil {
		domain = equipments.Domain{
			ID:				model.ID,
			CreatedAt:      model.CreatedAt,
			UpdatedAt:      model.UpdatedAt,
			Type: 			model.Type,
			Name:       	model.Name,
			Description: 	model.Description,
			ATK: 			model.ATK,
			DEF: 			model.DEF,
			HP: 			model.HP,	
		}
	}
	return domain
}
