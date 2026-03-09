package slack

import "context"

//go:generate mockgen -source=gateway.go -destination=./mocks/mock_gateway.go -package=mocks

type Gateway interface {
	SendMessage(ctx context.Context, message string) error
}