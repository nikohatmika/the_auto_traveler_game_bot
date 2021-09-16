package events

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
	Description string	
	GoldReward 	int		
	XPReward	int		
}

type Usecase interface {
	Find(ctx context.Context, eventType string) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, eventType string) ([]Domain, error)
}