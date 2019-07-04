package float

import (
	"testing"
)

func TestAvxFloat32Dot(t *testing.T) {
	x := []float32{1, 1, 1, 1, 1, 1, 1, 1}
	expected := float32(8)
	actual := avxFloat32Dot(x, x)
	if expected != actual {
		t.Errorf("expected %g actual %g", expected, actual)
	}
	x = append(x, 2)
	expected += 2 * 2
	actual = avxFloat32Dot(x, x)
	if expected != actual {
		t.Errorf("expected %g actual %g", expected, actual)
	}
	for i := 0; i < 16; i++ {
		x = append(x, x...)
		expected *= 2
		actual = avxFloat32Dot(x, x)
		if expected != actual {
			t.Errorf("expected %g actual %g", expected, actual)
		}
	}
}

func TestAvxFloat64Dot(t *testing.T) {
	x := []float64{1, 1, 1, 1, 1, 1, 1, 1}
	expected := float64(8)
	actual := avxFloat64Dot(x, x)
	if expected != actual {
		t.Errorf("expected %g actual %g", expected, actual)
	}
	x = append(x, 2)
	expected += 2 * 2
	actual = avxFloat64Dot(x, x)
	if expected != actual {
		t.Errorf("expected %g actual %g", expected, actual)
	}
	for i := 0; i < 16; i++ {
		x = append(x, x...)
		expected *= 2
		actual = avxFloat64Dot(x, x)
		if expected != actual {
			t.Errorf("expected %g actual %g", expected, actual)
		}
	}

}

func TestAvxFloat32Sum(t *testing.T) {
	x := []float32{1, 1, 1, 1, 1, 1, 1, 1}
	expected := float32(8)
	actual := avxFloat32Sum(x)
	if expected != actual {
		t.Errorf("expected %g actual %g", expected, actual)
	}
	x = append(x, 2)
	expected += 2
	actual = avxFloat32Sum(x)
	if expected != actual {
		t.Errorf("expected %g actual %g", expected, actual)
	}
	for i := 0; i < 16; i++ {
		x = append(x, x...)
		expected *= 2
		actual = avxFloat32Sum(x)
		if expected != actual {
			t.Errorf("expected %g actual %g", expected, actual)
		}
	}
}

func TestAvxFloat64Sum(t *testing.T) {
	x := []float64{1, 1, 1, 1, 1, 1, 1, 1}
	expected := float64(8)
	actual := avxFloat64Sum(x)
	if expected != actual {
		t.Errorf("expected %g actual %g", expected, actual)
	}
	x = append(x, 2)
	expected += 2
	actual = avxFloat64Sum(x)
	if expected != actual {
		t.Errorf("expected %g actual %g", expected, actual)
	}

}
