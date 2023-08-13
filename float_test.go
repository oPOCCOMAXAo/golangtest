package golangtest

import (
	"math"
	"strconv"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
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

func TestFloatToIntCastSaveInt(t *testing.T) {
	const maxSafeInt = 1<<53 - 1
	const minSafeInt = -(1<<53 - 1)
	t.Logf("max: %d, %b", maxSafeInt, maxSafeInt)
	t.Logf("min: %d, %b", minSafeInt, minSafeInt)

	floats := []float64{
		float64(math.MaxInt) * 1000,
		float64(math.MinInt) * 1000,
	}

	for _, floatVal := range floats {
		t.Logf("value: %f", floatVal)

		intVal := int(floatVal)
		t.Logf("raw cast: %d, %b", intVal, intVal)

		intVal = int(math.Round(floatVal))
		t.Logf("round cast: %d, %b", intVal, intVal)

		intVal = int(lo.Clamp(floatVal, minSafeInt, maxSafeInt))
		t.Logf("clamp cast: %d, %b", intVal, intVal)
		assert.Contains(t, []int{minSafeInt, maxSafeInt}, intVal)
	}
}

func TestFloatToIntCastOverflow(t *testing.T) {
	minInt := math.MinInt
	t.Logf("min: %d, %b", minInt, minInt)

	const maxSafeFloat = 1<<63 - 1<<9 - 1
	const minSafeFloat = -(1<<63 - 1<<9 - 1)
	t.Logf("max: %d, %b", maxSafeFloat, maxSafeFloat)
	t.Logf("min: %d, %b", minSafeFloat, minSafeFloat)

	floats := []float64{
		float64(math.MaxInt) * 1000,
		float64(math.MinInt) * 1000,
	}

	for _, floatVal := range floats {
		t.Logf("value: %f", floatVal)

		intVal := int(floatVal)
		t.Logf("raw cast: %d, %b", intVal, intVal)

		intVal = int(math.Round(floatVal))
		t.Logf("round cast: %d, %b", intVal, intVal)

		intVal = int(lo.Clamp(floatVal, minSafeFloat, maxSafeFloat))
		t.Logf("clamp cast: %d, %b", intVal, intVal)
		assert.NotEqual(t, math.MinInt, intVal)
	}
}
