package players

import (
	"context"
	"time"
)

type Domain struct {
	ID        			int
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	DeletedAt 			time.Time
	Email     			string
	Password  			string
	Name      			string
	Level				int
	XP					int
	Gold				int
	FamePoints			int
	SpareStatPoints		int
	ATK					int
	DEF					int
	HP					int
	WeaponEquipmentID	int
	ArmorEquipmentID	int
}

type Usecase interface {
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	Delete(ctx context.Context, ID int) (err error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	Delete(ctx context.Context, ID int) (err error)
}
