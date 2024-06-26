package ranges

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: write more and better tests.
func TestRanges(t *testing.T) {
	res := createRangeMap("./testdata/")
	array := strings.Split(res, "\n")
	assert.Equal(t, 38, len(array))
	m := GetRangeMap()
	assert.Equal(t, 36, len(m))
	w := "map[string][]int{\n"
	for k, v := range m {
		w += fmt.Sprintf("\"%s\": {%d, %d},\n", k, slices.Min(v), slices.Max(v))
	}
	w += "}"
}
