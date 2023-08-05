package golangtest

import (
	"math"
	"strconv"
	"testing"
)

func TestFloat(t *testing.T) {
	value := math.Nextafter(1, 0)
	t.Log(value)
	input := strconv.FormatFloat(value, 'f', -1, 64)
	t.Log(input)

	for a := 0; a < 200; a++ {
		var val float64
		for i := 9; i >= 0; i-- {
			temp := input + strconv.FormatInt(int64(i), 10)

			val, _ = strconv.ParseFloat(temp, 64)

			if val == value {
				input = temp
				break
			}
		}

		if val == 1 {
			break
		}
	}

	t.Log(input)
}
