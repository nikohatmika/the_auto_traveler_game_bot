package players

import (
	"auto_traveler/bussiness/players"
	"context"

	"gorm.io/gorm"
)

type playersRepository struct {
	conn *gorm.DB
}

func NewPlayersRepository(conn *gorm.DB) players.Repository {
	return &playersRepository{
		conn: conn,
	}
}

func (r *playersRepository) FindByID(ctx context.Context, ID int) (players.Domain, error) {
	var res *Players
	err := r.conn.Where("deleted_at", nil).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return players.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *playersRepository) FindByEmail(ctx context.Context, email string) (players.Domain, error) {
	var res *Players
	err := r.conn.Where("deleted_at", nil).Where("email = ?", email).First(&res).Error
	if err != nil {
		return players.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *playersRepository) Store(ctx context.Context, data *players.Domain) (res players.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Create(&model)
	if result.Error != nil {
		return players.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *playersRepository) Update(ctx context.Context, ID int, data *players.Domain) (res players.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Where("id = ?", ID).Save(&model)
	if result.Error != nil {
		return players.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

