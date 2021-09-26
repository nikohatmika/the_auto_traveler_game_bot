package players

import (
	"auto_traveler/helper/encrypt"
	"context"
	"time"
)

type PlayerUsecase struct {
	playerRespository 	Repository
	contextTimeout      time.Duration
}

func NewPlayerUsecase(timeout time.Duration, r Repository) Usecase {
	return &PlayerUsecase{
		contextTimeout:     timeout,
		playerRespository: 	r,
	}
}


func (uc *PlayerUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := uc.playerRespository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *PlayerUsecase) FindByEmail(ctx context.Context, email string) (Domain, error) {
	resp, err := uc.playerRespository.FindByEmail(ctx, email)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}
func (uc *PlayerUsecase) Store(ctx context.Context, data *Domain) (res Domain, err error) {
	data.Password, err = encrypt.Hash(data.Password)
	if err != nil {
		return res, err
	}

	res, err = uc.playerRespository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

