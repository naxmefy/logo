package logo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestI struct {
	ID    int
	Title string
}

var testi1 = TestI{
	ID:    1,
	Title: "Test 1",
}
var testi2 = TestI{
	ID:    2,
	Title: "Test 2",
}
var testi3 = TestI{
	ID:    3,
	Title: "Test 3",
}
var testi4 = TestI{
	ID:    4,
	Title: "Test 4",
}

var testis = []TestI{
	testi1,
	testi2,
	testi3,
	testi4,
}

func TestIndex(t *testing.T) {
	assert.Equal(t, Index(testis, testi1), 0)
	assert.Equal(t, Index(testis, testi2), 1)
	assert.Equal(t, Index(testis, testi3), 2)
	assert.Equal(t, Index(testis, testi4), 3)
	assert.Equal(t, Index(testis, TestI{}), -1)

	assert.Equal(t, Index("test", "t"), 0)
	assert.Equal(t, Index("test", "e"), 1)
	assert.Equal(t, Index("test", "s"), 2)

	assert.Equal(t, Index([]string{"hello", "world"}, "world"), 1)
}

func TestInclude(t *testing.T) {
	assert.True(t, Include(testis, testi1))
	assert.True(t, Include(testis, testi2))
	assert.True(t, Include(testis, testi3))
	assert.True(t, Include(testis, testi4))
	assert.False(t, Include(testis, TestI{}))

	assert.True(t, Include("test", "t"))
	assert.True(t, Include("test", "tes"))
	assert.False(t, Include("test", "a"))

	assert.True(t, Include([]string{"hello", "world"}, "world"))
}

func TestAny(t *testing.T) {
	assert.True(t, Any(testis, func(v interface{}) bool {
		return v.(TestI).ID == 1
	}))

	assert.True(t, Any("test", func(v interface{}) bool {
		return strings.Contains(v.(string), "t")
	}))

	assert.False(t, Any("test", func(v interface{}) bool {
		return strings.Contains(v.(string), "a")
	}))
}

func TestAll(t *testing.T) {
	assert.True(t, All(testis, func(v interface{}) bool {
		return strings.Contains(v.(TestI).Title, "Test")
	}))

	assert.True(t, All("tttt", func(v interface{}) bool {
		return strings.Contains(v.(string), "t")
	}))

	assert.False(t, All("test", func(v interface{}) bool {
		return strings.Contains(v.(string), "t")
	}))
}

func TestFilter(t *testing.T) {
	t.Skip()
}

func TestMap(t *testing.T) {
	t.Skip()
}
