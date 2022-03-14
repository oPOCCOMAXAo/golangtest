package golangtest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonUnmarshalToNil(t *testing.T) {
	assert.Error(t, json.Unmarshal([]byte("{}"), nil))
}
