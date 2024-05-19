package distance

import (
	"testing"
)

func TestDistance(t *testing.T) {
	t.Run("when given two coordinates give integer as answer", func(t *testing.T) {
		p1 := NewPoint(1, 1)
		p2 := NewPoint(4, 5)
		if d := Distance(p1, p2); d != 5 {
			t.Errorf("expected  5  but got %v", d)
		}
	})
	t.Run("when given two coordinates do not give integer as answer", func(t *testing.T) {
		p1 := NewPoint(1, 1)
		p2 := NewPoint(3, 5)
		if d := Distance(p1, p2); d != 4.47213595499958 {
			t.Errorf("expected  4.47213595499958  but got %v", d)
		}
	})
}
