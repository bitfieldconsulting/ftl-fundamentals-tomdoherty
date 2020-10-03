package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b    float64
		want    float64
		comment string
	}{
		{a: 1, b: 2, want: 3, comment: "one add two"},
		{a: 2, b: 5, want: 7, comment: "two add five"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64() + 1
		b := rand.Float64() + 1
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b    float64
		want    float64
		comment string
	}{
		{a: 1, b: 2, want: -1, comment: "one minus two"},
		{a: 2, b: 5, want: -3, comment: "two minus five"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, tc.want, got)
		}
	}
}

func TestSubtractRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64() + 1
		b := rand.Float64() + 1
		want := a - b
		got := calculator.Subtract(a, b)
		if want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b    float64
		want    float64
		comment string
	}{
		{a: 1, b: 2, want: 2, comment: "one times two"},
		{a: 2, b: 5, want: 10, comment: "two times five"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, tc.want, got)
		}
	}
}

func TestMultiplyRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64() + 1
		b := rand.Float64() + 1
		want := a * b
		got := calculator.Multiply(a, b)
		if want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b        float64
		want        float64
		comment     string
		errExpected bool
	}{
		{a: 2, b: 1, want: 2, errExpected: false, comment: "two divided by 1"},
		{a: 10, b: 5, want: 2, errExpected: false, comment: "ten divided by 5"},
		{a: 10, b: 0, want: 0, errExpected: true, comment: "error out on divide by zero"},
		{a: 10, b: 0, want: 1, errExpected: true, comment: "test error case by failing want vs get"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errorReceived := err != nil
		if errorReceived != tc.errExpected {
			t.Fatalf("test %s: expected err %t, got %t", tc.comment, tc.errExpected, errorReceived)
		}
		if !errorReceived && tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, tc.want, got)
		}
	}
}

func TestDivideRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64() + 1
		b := rand.Float64() + 1
		want := a / b
		got, err := calculator.Divide(a, b)
		if err != nil {
			t.Errorf("Divide(%f, %f): want %f, got %f", a, b, want, got)
		}
		if want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		a, b        float64
		want        float64
		comment     string
		errExpected bool
	}{
		{a: 16, want: 4, errExpected: false, comment: "square root of 16"},
		{a: -1, want: 0, errExpected: true, comment: "error"},
		{a: -1, want: 0, errExpected: true, comment: "error out on negative input"},
		{a: -1, want: 1, errExpected: true, comment: "test error case by failing want vs get"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errorReceived := err != nil
		if errorReceived != tc.errExpected {
			t.Fatalf("test %s: expected err %t, got %t", tc.comment, tc.errExpected, errorReceived)
		}
		if !errorReceived && tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, tc.want, got)
		}
	}
}

func TestSqrtRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64() + 1
		want := math.Sqrt(a)
		got, err := calculator.Sqrt(a)
		if err != nil {
			t.Errorf("Sqrt(%f): want %f, got %f", a, want, got)
		}
		if want != got {
			t.Errorf("Sqrt(%f): random test: want %f, got %f", a, want, got)
		}
	}
}
