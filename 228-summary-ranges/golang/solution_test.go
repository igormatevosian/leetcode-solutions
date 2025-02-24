package golang

import (
	"slices"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []string{"0->2", "4->5", "7"}
	actual := summaryRanges([]int{0, 1, 2, 4, 5, 7})

	if !slices.Equal(expected, actual) {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := []string{"0", "2->4", "6", "8->9"}
	actual := summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})

	if !slices.Equal(expected, actual) {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}

}
