package golangtest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type jsonAny struct {
	A int64
	B float64
	C string
	D bool
}

func (r *jsonAny) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &[]interface{}{&r.A, &r.B, &r.C, &r.D})
}

func TestJSONMixedArray(t *testing.T) {
	input := `[1, 2, "3", true]`
	var res jsonAny

	assert.NoError(t, json.Unmarshal([]byte(input), &res))
	assert.Equal(t, jsonAny{
		A: 1,
		B: 2,
		C: "3",
		D: true,
	}, res)
}
