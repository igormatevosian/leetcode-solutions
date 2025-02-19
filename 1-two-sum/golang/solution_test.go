package golang

import (
	"reflect"
	"slices"
	"testing"
)

func TestCase1(t *testing.T) {
	input := struct {
		nums   []int
		target int
	}{
		nums:   []int{2, 7, 11, 15},
		target: 9,
	}
	expected := []int{0, 1}

	actual := twoSum(input.nums, input.target)
	revActual := actual
	slices.Reverse(revActual)
	if !reflect.DeepEqual(expected, actual) || !reflect.DeepEqual(expected, revActual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	input := struct {
		nums   []int
		target int
	}{
		nums:   []int{3, 2, 4},
		target: 6,
	}
	expected := []int{1, 2}

	actual := twoSum(input.nums, input.target)
	revActual := actual
	slices.Reverse(revActual)
	if !reflect.DeepEqual(expected, actual) || !reflect.DeepEqual(expected, revActual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	input := struct {
		nums   []int
		target int
	}{
		nums:   []int{3, 3},
		target: 6,
	}
	expected := []int{0, 1}

	actual := twoSum(input.nums, input.target)
	revActual := actual
	slices.Reverse(revActual)
	if !reflect.DeepEqual(expected, actual) || !reflect.DeepEqual(expected, revActual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}
