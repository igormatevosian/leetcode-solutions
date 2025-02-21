package golang

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expected := 2

	actual := findMedianSortedArrays(nums1, nums2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	expected := 2.5

	actual := findMedianSortedArrays(nums1, nums2)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}
