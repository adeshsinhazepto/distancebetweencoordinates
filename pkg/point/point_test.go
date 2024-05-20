package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Point(t *testing.T) {
	t.Run("should be able to create new point given two coordinates", func(t *testing.T) {
		assert.NotPanics(t, func() { NewPoint(3, 4) })
	})
	t.Run("should be able to check returning correct points", func(t *testing.T) {
		p := NewPoint(3, 4)
		assert.Equal(t, 3, p.X)
		assert.Equal(t, 4, p.Y)
	})
}
