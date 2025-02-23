package golang

import "testing"

func TestCase1(t *testing.T) {
	expected := 2
	actual := maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1})

	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := 3
	actual := maxDistToClosest([]int{1, 0, 0, 0})

	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}
