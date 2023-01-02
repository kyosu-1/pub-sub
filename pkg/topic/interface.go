package topic

import (
	"context"

	"gocloud.dev/pubsub"
)

type Interface interface {
	Send(ctx context.Context, m *pubsub.Message) (err error)
}