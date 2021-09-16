package equipments

import (
	"context"
	"time"
)

type Domain struct {
	ID 			int
	CreatedAt	time.Time
	UpdatedAt 	time.Time		
	Type		string	
	Name 		string	
	Description	string	
	ATK 		int		
	DEF			int
	HP 			int		
}

type Usecase interface {
	Find(ctx context.Context, eqType string) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, eqType string) ([]Domain, error)
}