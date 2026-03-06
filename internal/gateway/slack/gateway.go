package slack

import "context"

type Gateway interface {
	SendMessage(ctx context.Context, message string) error
}