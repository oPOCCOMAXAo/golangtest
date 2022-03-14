package golangtest

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestLevel3(t *testing.T) {
	println("Location:", gpsRequest(newAgent()))
}

func newAgent() string {
	return "a a a"
}

func gpsRequest(agent string) string {
	a := strings.Split(agent, " ")

	denied := len(a) != 3 ||
		!strings.EqualFold(a[0], a[1]) ||
		!(strings.Count(a[2], a[1]) > 0)

	if denied {
		return "ACCESS DENIED"
	}

	return "\u0031\u0036\u002e\u0037\u0033\u0033\u0033\u002c\u002d\u0031\u0036\u0039\u002e\u0035\u0032\u0037\u0034"
}

func TestLevel8(t *testing.T) {
	println("Logging in...")
	authorized := startup(login())
	if reflect.ValueOf(authorized).Bool() {
		println("Starting the engine")
		return
	}
	println("Startup failed")
}

func validSequence(i int, el interface{}) bool {
	fmt.Printf("%d %#v\n", i, el)
	println(
		reflect.TypeOf(el).String() == "*golangtest.Sequence",
		!reflect.ValueOf(el).IsNil(),
		reflect.ValueOf(el).Elem().NumField() == 2,
		reflect.TypeOf(reflect.ValueOf(el).Elem().Field(0).Interface()).String(),
		int(reflect.ValueOf(el).Elem().Field(0).Int()) == i*i-i,
		!reflect.ValueOf(reflect.ValueOf(el).Elem().Field(1).Interface()).IsNil(),
	)
	return reflect.TypeOf(el).String() == "*golangtest.Sequence" &&
		!reflect.ValueOf(el).IsNil() &&
		reflect.ValueOf(el).Elem().NumField() == 2 &&
		reflect.TypeOf(reflect.ValueOf(el).Elem().Field(0).Interface()).String() == "int" &&
		int(reflect.ValueOf(el).Elem().Field(0).Int()) == i*i-i &&
		!reflect.ValueOf(reflect.ValueOf(el).Elem().Field(1).Interface()).IsNil()
}

func startup(seq interface{}) bool {
	for i := 0; i < 5; i++ {
		if !validSequence(i, seq) {
			return false
		}
		seq = reflect.ValueOf(seq).Elem().Field(1).Interface()
	}

	return true
}

type Sequence struct {
	Field1 int
	Field2 *Sequence
}

func login() *Sequence {
	return &Sequence{
		Field1: 0,
		Field2: &Sequence{
			Field1: 0,
			Field2: &Sequence{
				Field1: 2,
				Field2: &Sequence{
					Field1: 6,
					Field2: &Sequence{
						Field1: 12,
						Field2: &Sequence{
							Field1: 20,
						},
					},
				},
			},
		},
	}
}
