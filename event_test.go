package golangtest

import (
	"testing"

	"github.com/opoccomaxao-go/event/v2"
)

func TestEvent(t *testing.T) {
	pool := event.NewPool(event.PoolConfig{})

	_ = pool
}
