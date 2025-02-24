package golang

import "testing"

func TestCase1(t *testing.T) {
	expected := 3
	actual := longestSubarray([]int{1, 1, 0, 1})

	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := 5
	actual := longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})

	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := 2
	actual := longestSubarray([]int{1, 1, 1})

	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}
