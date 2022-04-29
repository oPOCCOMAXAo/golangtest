package golangtest

import (
	"testing"

	"github.com/opoccomaxao-go/event/v3"
)

func TestEvent(t *testing.T) {
	pool := event.NewPool(event.PoolConfig{})

	_ = pool
}
