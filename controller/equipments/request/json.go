package request

import (
	"auto_traveler/bussiness/equipments"
	"time"
)

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

func (req *Equipments) ToDomain() (res *equipments.Domain, err error) {
	return &equipments.Domain{
		ID:				req.ID,		 		
		CreatedAt:  	time.Now().UTC(),
		UpdatedAt:  	time.Now().UTC(),
		Type: 			req.Type,
		Name:			req.Name,
		Description: 	req.Description, 
		ATK: 			req.ATK,
		DEF: 			req.DEF,
		HP: 			req.HP,
	}, err
}