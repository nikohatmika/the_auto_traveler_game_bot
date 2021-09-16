package equipments

import (
	"auto_traveler/bussiness/equipments"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type equipmentsRepository struct {
	conn *gorm.DB
}

func NewEquipmentsRepository(conn *gorm.DB) equipments.Repository {
	return &equipmentsRepository{
		conn: conn,
	}
}

func (r *equipmentsRepository) Find(ctx context.Context, eqType string) ([]equipments.Domain, error) {
	res := []Equipments{}

	query := r.conn

	if eqType != "" {
		query = query.Where("type", eqType)
	}

	err := query.Find(&res).Error
	if err != nil {
		return []equipments.Domain{}, err
	}

	equipmentsDomain := []equipments.Domain{}
	for _, value := range res {
		equipmentsDomain = append(equipmentsDomain, value.ToDomain())
	}

	fmt.Println("res", res)
	// fmt.Println("query", query)

	return equipmentsDomain, nil
}