package finalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsStringHelloSonartwo(t *testing.T) {
	e := ContainsStringHelloSonartwo([]string{"a", "b"}, "b")
	assert.True(t, e)
}

