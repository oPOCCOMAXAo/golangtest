package golangtest

import (
	"fmt"
	"reflect"
	"testing"
)

func typef(typ reflect.Type) string {
	return fmt.Sprintf("Name: %s\nAlign: %v\nComparable: %v\nKind: %#x\nSize: %d\nString: %s\n",
		typ.Name(),
		typ.Align(),
		typ.Comparable(),
		typ.Kind(),
		typ.Size(),
		typ.String(),
	)
}

func TestType(t *testing.T) {
	type t1 struct{}
	type t2 struct{}
	type t3 = t1

	t.Logf("t1 %s\n", typef(reflect.TypeOf(t1{})))
	t.Logf("t2 %s\n", typef(reflect.TypeOf(t2{})))
	t.Logf("t3 %s\n", typef(reflect.TypeOf(t3{})))
	t.Logf("*t1 %s\n", typef(reflect.TypeOf(&t1{})))
	t.Logf("*t2 %s\n", typef(reflect.TypeOf(&t2{})))
	t.Logf("*t3 %s\n", typef(reflect.TypeOf(&t3{})))
}
