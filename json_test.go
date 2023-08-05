package golangtest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJsonUnmarshalToNil(t *testing.T) {
	assert.Error(t, json.Unmarshal([]byte("{}"), nil))
}

func TestJsonByteArray(t *testing.T) {
	arr := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	data, err := json.Marshal(arr)
	require.NoError(t, err)

	assert.Equal(t, `"AQIDBAUGBwg="`, string(data))
}

func TestJSONStringEmpty(t *testing.T) {
	data, err := json.Marshal(struct {
		A string `json:"a,omitempty"`
		B string `json:"b,omitempty"`
	}{B: "1"})
	require.NoError(t, err)
	assert.Equal(t, `{"b":"1"}`, string(data))
}
