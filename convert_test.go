package struct2mapstrstr

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestConvertMapStringString(t *testing.T) {
	type MssCheck struct {
		Test string  `mss:"test"`
		Tas  string  `mss:"tassss"`
		Nas  int     `mss:"nas"`
		X    *string `mss:"x"`
	}

	x := "x"
	v := MssCheck{
		Test: "iaiu",
		Tas:  "ooo",
		Nas:  3,
		X:    &x,
	}
	k := ConvertMapStringString(v)
	assert.Equal(t, map[string]string{"test": "iaiu", "tassss": "ooo", "nas": "3", "x": "x"}, k)

	v = MssCheck{
		Test: "iaiu",
		Tas:  "ooo",
		Nas:  3,
	}
	k = ConvertMapStringString(v)
	assert.Equal(t, map[string]string{"test": "iaiu", "tassss": "ooo", "nas": "3"}, k)
}
