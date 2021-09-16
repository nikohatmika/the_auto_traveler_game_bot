package player_auth

import (
	"auto_traveler/app/middleware"
	"auto_traveler/bussiness/players"
	"auto_traveler/helper/encrypt"
	"auto_traveler/helper/messages"
	"context"
	"strings"
	"time"

	"gorm.io/gorm"
)

type playerAuthUsecase struct {
	playerRepository 	players.Repository
	contextTimeout     	time.Duration
	jwtAuth            	*middleware.ConfigJWT
}

func NewplayerAuthUsecase(timeout time.Duration, playerRepo players.Repository, jwt *middleware.ConfigJWT) Usecase {
	return &playerAuthUsecase{
		playerRepository: playerRepo,
		jwtAuth:            jwt,
		contextTimeout:     timeout,
	}
}

func (uc playerAuthUsecase) Register(ctx context.Context, data *players.Domain) (res Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	player, err := uc.playerRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return res, err
		}
	}

	if player.ID != 0 {
		return res, messages.ErrDataAlreadyExist
	}

	data.Password, err = encrypt.Hash(data.Password)
	if err != nil {
		return res, err
	}

	player, err = uc.playerRepository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	res.Token = uc.jwtAuth.GenerateToken(player.ID, "player")

	return res, err
}

func (uc playerAuthUsecase) Login(ctx context.Context, data *Domain) (res Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(data.Email) == "" && strings.TrimSpace(data.Password) == "" {
		return res, messages.ErrUsernamePasswordNotFound
	}

	player, err := uc.playerRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return res, err
	}

	if !encrypt.ValidateHash(data.Password, player.Password) {
		return res, messages.ErrInvalidCred
	}

	res.Token = uc.jwtAuth.GenerateToken(player.ID, "player")

	return res, err
}
