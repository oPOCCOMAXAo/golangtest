package golangtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChan_closeClosed(t *testing.T) {
	ch := make(chan struct{})
	close(ch)

	assert.Panics(t, func() {
		close(ch)
	})
}

func TestChan_closeNil(t *testing.T) {
	var ch chan struct{}

	assert.Panics(t, func() {
		close(ch)
	})
}
