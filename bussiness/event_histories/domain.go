package event_histories

import (
	"context"
	"time"
)

type Domain struct {
	ID 			int
	CreatedAt	time.Time
	PlayerID	int
	EventID		int
	Result		bool	
}

type Usecase interface {
	Find(ctx context.Context, playerId int) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, playerId int) ([]Domain, error)
}