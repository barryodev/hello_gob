package hello_gob

import (
	"testing"
)

func TestConvertingPointsToBytes(t *testing.T) {
	newPoint := Point{X: 10, Y: 15}

	bytesArray := newPoint.PointToBytes()

	nextPoint, err := BytesToPoint(bytesArray)

	if err != nil {
		t.Errorf("got error")
	}

	if newPoint.X != nextPoint.X {
		t.Errorf("got %q, wanted %q", newPoint.X,  nextPoint.X)
	}

	if newPoint.Y != nextPoint.Y {
		t.Errorf("got %q, wanted %q", newPoint.Y,  nextPoint.Y)
	}
}

func TestHandlesEmptyByteArray(t *testing.T) {
	newPoint, err := BytesToPoint(nil)

	if err == nil {
		t.Errorf("expected error")
	}

	if err.Error() != "Failed to decode bytes: EOF" {
		t.Errorf("got %q, wanted %q", err.Error(),  "Failed to decode bytes: EOF")
	}

	if newPoint.X != 0 {
		t.Errorf("got %q, wanted %q", newPoint.X,  0)
	}

	if newPoint.Y != 0 {
		t.Errorf("got %q, wanted %q", newPoint.Y,  0)
	}
}

