package players

import (
	"auto_traveler/bussiness/players"
	"database/sql"
	"time"
)

type Players struct {
	ID        int
	Name      string
	Email     string
	Password  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func fromDomain(domain *players.Domain) *Players {
	return &Players{
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  sql.NullString{String: domain.Password, Valid: true},
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (model *Players) ToDomain() (domain *players.Domain) {
	if model != nil {
		domain = &players.Domain{
			ID:        model.ID,
			Name:      model.Name,
			Email:     model.Email,
			Password:  model.Password.String,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt.Time,
		}
	}
	return domain
}
