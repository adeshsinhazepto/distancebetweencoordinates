package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Point(t *testing.T) {
	t.Run("should be able to create new point given two coordinates", func(t *testing.T) {
		assert.NotPanics(t, func() { NewPoint(3, 4) })
	})
}
