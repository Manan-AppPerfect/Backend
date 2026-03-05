package user

import "context"

type Gateway interface {
	GetUser(ctx context.Context, id int64) (string, error)
}