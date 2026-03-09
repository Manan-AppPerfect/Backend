package user

import "context"

//go:generate mockgen -source=gateway.go -destination=./mocks/mock_gateway.go -package=mocks

type Gateway interface {
	GetUser(ctx context.Context, id int64) (string, error)
}