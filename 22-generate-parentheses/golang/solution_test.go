package golang

import (
	"slices"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	actual := generateParenthesis(3)

	if !slices.Equal(expected, actual) {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := []string{"()"}
	actual := generateParenthesis(1)

	if !slices.Equal(expected, actual) {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}

}
