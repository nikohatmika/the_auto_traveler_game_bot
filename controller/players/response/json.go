package response

import "auto_traveler/bussiness/players"

type Player struct {
	ID        			int			`json:"id"`
	CreatedAt 			string	`json:"created_at"`
	UpdatedAt 			string	`json:"updated_at"`
	DeletedAt 			string	`json:"deleted_at"`
	Email     			string		`json:"email"`
	Password  			string		`json:"password"`
	Name      			string		`json:"name"`
	Level				int			`json:"level"`
	XP					int			`json:"xp"`
	Gold				int			`json:"gold"`
	FamePoints			int			`json:"fame_points"`
	SpareStatPoints		int			`json:"spare_stat_points"`
	ATK					int			`json:"atk"`
	DEF					int			`json:"def"`
	HP					int			`json:"hp"`
	WeaponEquipmentID	int			`json:"weapon_equipment_id"`
	ArmorEquipmentID	int			`json:"armor_equipment_id"`
}

func FromDomain(domain *players.Domain) (res *Player) {
	if domain != nil {
		res = &Player{
			ID:        domain.ID,
			Name:      domain.Name,
			Email:     domain.Email,
			CreatedAt: domain.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
			UpdatedAt: domain.UpdatedAt.UTC().Format("2006-01-02 15:04:05"),
			DeletedAt: domain.DeletedAt.UTC().Format("2006-01-02 15:04:05"),
		}
	}

	return res
}
