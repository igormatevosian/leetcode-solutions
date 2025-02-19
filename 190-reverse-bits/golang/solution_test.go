package golang

import "testing"

func TestCase1(t *testing.T) {
	expected := uint32(964176192)
	actual := reverseBits(43261596)

	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := uint32(3221225471)
	actual := reverseBits(4294967293)

	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}
