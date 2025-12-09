package main

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	lines := []string{
		"L1",
		"L50",
		"L99",
		"L100",
		"R1",
		"R50",
		"R99",
		"R100",
	}

	expected := []int{
		-1,
		-50,
		-99,
		0,
		1,
		50,
		99,
		0,
	}

	actual := Convert(lines)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Convert() = %v, want %v", actual, expected)
	}
}

func TestComputeResult(t *testing.T) {
	lines := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	expected := 6
	actual, _ := ComputeResult(lines)

	if actual != expected {
		t.Errorf("ComputeResult() = %v, want %v", actual, expected)
	}
}
