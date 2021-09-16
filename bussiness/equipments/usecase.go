package equipments

import (
	"context"
	"time"
)

type equipmentsUsecase struct {
	equipmentsRepository Repository
	contextTimeout time.Duration
}

func NewEquipmentsUsecase(timeout time.Duration, r Repository) Usecase {
	return &equipmentsUsecase{
		contextTimeout: timeout,
		equipmentsRepository: r,
	}
}

func (uc *equipmentsUsecase) Find(ctx context.Context) ([]Domain, error) {
	resp, err := uc.equipmentsRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}