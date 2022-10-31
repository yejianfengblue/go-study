package string

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLastIndex(t *testing.T) {
	assert.Equal(t, 4, strings.LastIndex("1234 567890", " "))
}

func TestToUpper(t *testing.T) {
	assert.Equal(t, "A", strings.ToUpper("a"))
}
