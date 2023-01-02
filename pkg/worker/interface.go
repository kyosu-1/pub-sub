package worker

import (
	"context"

	"gocloud.dev/pubsub"
)

type Interface interface {
	Run(ctx context.Context)
	Handle(ctx context.Context, msg *pubsub.Message)
}