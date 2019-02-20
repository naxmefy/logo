package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var strs = []string{"peach", "apple", "pear", "plum"}

func TestIndex(t *testing.T) {
	assert.Equal(t, Index(strs, "peach"), 0)
	assert.Equal(t, Index(strs, "apple"), 1)
	assert.Equal(t, Index(strs, "pear"), 2)
	assert.Equal(t, Index(strs, "plum"), 3)
	assert.Equal(t, Index(strs, "somethingelse"), -1)
}

func TestInclude(t *testing.T) {
	assert.Equal(t, Include(strs, "peach"), true)
	assert.Equal(t, Include(strs, "apple"), true)
	assert.Equal(t, Include(strs, "pear"), true)
	assert.Equal(t, Include(strs, "plum"), true)
	assert.Equal(t, Include(strs, "somethingelse"), false)
}

func TestAny(t *testing.T) {
	assert.Equal(t, Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}), true)
	assert.Equal(t, Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "s")
	}), false)
}

func TestAll(t *testing.T) {
	assert.Equal(t, All(strs, func(v string) bool {
		return strings.ContainsAny(v, "p")
	}), true)
	assert.Equal(t, All(strs, func(v string) bool {
		return strings.ContainsAny(v, "s")
	}), false)
}

func TestFilter(t *testing.T) {
	someStrs := Filter(strs, func(v string) bool {
		return strings.Contains(v, "apple")
	})

	assert.Contains(t, someStrs, "apple")
	assert.NotContains(t, someStrs, "peach")
	assert.NotContains(t, someStrs, "pear")
	assert.NotContains(t, someStrs, "plum")
}

func TestMap(t *testing.T) {
	someStrs := Map(strs, strings.ToUpper)
	assert.Contains(t, someStrs, "APPLE")
	assert.Contains(t, someStrs, "PEACH")
	assert.Contains(t, someStrs, "PEAR")
	assert.Contains(t, someStrs, "PLUM")

	assert.NotContains(t, someStrs, "apple")
	assert.NotContains(t, someStrs, "peach")
	assert.NotContains(t, someStrs, "pear")
	assert.NotContains(t, someStrs, "plum")
}
