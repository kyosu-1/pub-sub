package subscription

import (
	"context"

	"gocloud.dev/pubsub"
)

type Interface interface {
	Receive(ctx context.Context) (_ *pubsub.Message, err error)
}