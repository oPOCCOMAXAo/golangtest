package golangtest

import (
	"testing"

	"github.com/opoccomaxao-go/reactive-vars/reactive"
)

func TestReactive(t *testing.T) {
	type test struct {
		A float64
		B float64
	}

	v := reactive.NewSyncVariable[test]("test")
	v.OnChange(func(v test) {
		t.Logf("%+v\n", v)
	})

	v.Set(test{})
	v.Set(test{
		A: 1,
	})
	v.Set(test{
		A: 1,
	})
	v.Set(test{
		B: 1,
	})
	v.Set(test{
		B: 1,
	})

}
