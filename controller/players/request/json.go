package request

import "auto_traveler/bussiness/players"


type Player struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *Player) ToDomain() *players.Domain {
	return &players.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
