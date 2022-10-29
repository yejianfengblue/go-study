package math

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRound(t *testing.T) {

	ass := assert.New(t)

	ass.EqualValues(1, math.Round(1.2))
	ass.EqualValues(2, math.Round(1.5))
	ass.EqualValues(-1, math.Round(-1.2))
	ass.EqualValues(-2, math.Round(-1.5))

	ass.EqualValues(2, math.Ceil(1.2))
	ass.EqualValues(2, math.Ceil(1.5))
	ass.EqualValues(-1, math.Ceil(-1.2))
	ass.EqualValues(-1, math.Ceil(-1.5))

	ass.EqualValues(1, math.Floor(1.2))
	ass.EqualValues(1, math.Floor(1.5))
	ass.EqualValues(-2, math.Floor(-1.2))
	ass.EqualValues(-2, math.Floor(-1.5))
}
