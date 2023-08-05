package golangtest

import (
	"context"
	"testing"
)

func TestCancelCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	cancel()
	cancel()
	cancel()
	cancel()

	<-ctx.Done()
}
