package equipments

import (
	"context"
	"time"
)

type Domain struct {
	ID 			int
	CreatedAt	time.Time
	UpdatedAt 	time.Time		
	TypeID		int		
	Name 		string	
	Description	string	
	Atk 		int		
	Def			int
	HP 			int		
}

type Usecase interface {
	Find(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
}