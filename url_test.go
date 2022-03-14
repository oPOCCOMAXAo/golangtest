package golangtest

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	val := url.Values{
		"empty":  {},
		"nil":    nil,
		"single": {"1"},
		"double": {"1", "2"},
	}

	assert.Equal(t, "double=1&double=2&single=1", val.Encode())
}
