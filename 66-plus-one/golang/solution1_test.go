package golang

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 4}

	actual := plusOne(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	input := []int{4, 3, 2, 1}
	expected := []int{4, 3, 2, 2}

	actual := plusOne(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	input := []int{9}
	expected := []int{1, 0}

	actual := plusOne(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}
