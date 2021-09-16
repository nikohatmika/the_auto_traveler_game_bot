package player_auth

import (
	"auto_traveler/bussiness/players"
	"context"
)

type Domain struct {
	Email    string
	Password string
	Token    string
}

type Usecase interface {
	Login(ctx context.Context, data *Domain) (res Domain, err error)
	Register(ctx context.Context, data *players.Domain) (res Domain, err error)
}
