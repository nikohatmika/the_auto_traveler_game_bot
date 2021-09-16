package response

import "auto_traveler/bussiness/equipments"

type Equipments struct {
	ID 			int 	`json:"id"`
	CreatedAt	string 	`json:"created_at"`
	UpdatedAt 	string	`json:"updated_at"`	
	Type		string	`json:"type"`
	Name 		string	`json:"name"`
	Description string	`json:"description"`
	ATK 		int		`json:"atk"`	
	DEF			int		`json:"def"`
	HP			int		`json:"hp"`
}

func FromDomain(domain *equipments.Domain) (res *Equipments) {
	if domain != nil {
		res = &Equipments{
			ID:      		domain.ID,
			CreatedAt:      domain.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
			UpdatedAt:      domain.UpdatedAt.UTC().Format("2006-01-02 15:04:05"),
			Type: 			domain.Type,
			Name:         	domain.Name,
			Description:  	domain.Description,
			ATK: 			domain.ATK,
			DEF: 			domain.DEF,
			HP: 			domain.HP,
		}
	}

	return res
}