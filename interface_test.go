package golangtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceCast(t *testing.T) {
	type temp struct {
		A int
	}

	var val interface{} = &temp{A: 1}

	tval, ok := val.(*temp)
	assert.True(t, ok)
	assert.Equal(t, &temp{A: 1}, tval)

	i, ok := val.(*int)
	assert.False(t, ok)
	assert.Equal(t, (*int)(nil), i)

	var f *float64
	assert.Panics(t, func() {
		f = val.(*float64)
	})
	assert.Equal(t, (*float64)(nil), f)
}
