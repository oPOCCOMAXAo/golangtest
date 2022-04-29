package golangtest

import (
	"bytes"
	e0 "errors"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

var ErrTest = e0.New("test")

func TestErrorWrap(t *testing.T) {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%+v\n", ErrTest)
	fmt.Printf("%s\n", buffer.String())
	buffer.Reset()

	fmt.Fprintf(&buffer, "%+v\n", errors.WithStack(ErrTest))
	fmt.Printf("%s\n", buffer.String())
	buffer.Reset()

	fmt.Fprintf(&buffer, "%+v\n", errors.Wrap(ErrTest, "1"))
	fmt.Printf("%s\n", buffer.String())
	buffer.Reset()

	fmt.Fprintf(&buffer, "%+v\n", errors.WithMessage(ErrTest, "2"))
	fmt.Printf("%s\n", buffer.String())
	buffer.Reset()

}
