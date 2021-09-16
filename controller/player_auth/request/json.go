package request

import (
	"auto_traveler/bussiness/player_auth"
)

type PlayerAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (req *PlayerAuth) ToDomain() *player_auth.Domain {
	return &player_auth.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
