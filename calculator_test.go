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
		test string
	}{
		{a: 1, b: 2, want: 3, test: "one add two"},
		{a: 2, b: 5, want: 7, test: "two add five"},
		{a: -2, b: 2, want: 0, test: "minus two add two"},
		{a: -2, b: -2, want: -4, test: "minus two add minus two"},
		{a: -0, b: -2, want: -2, test: "zero add minus two"},
		{a: -2.6, b: 0.4, want: -2.2, test: "minus two point six add zero point four"},
		{a: math.Inf(1), b: -1, want: math.Inf(1), test: "+Inf add minus one"},
		{a: math.Inf(-1), b: -1, want: math.Inf(-1), test: "-Inf add minus one"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.test, tc.want, got)
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
		test string
	}{
		{a: 1, b: 2, want: -1, test: "one minus two"},
		{a: 2, b: 5, want: -3, test: "two minus five"},
		{a: -2, b: 2, want: -4, test: "minus two minus two"},
		{a: -2, b: -2, want: 0, test: "minus two minus minus two"},
		{a: 0, b: -2, want: 2, test: "zero minus minus two"},
		{a: -2.6, b: 0.4, want: -3, test: "minus two point six minus zero point four"},
		{a: math.Inf(1), b: -1, want: math.Inf(1), test: "+Inf minus minus one"},
		{a: math.Inf(-1), b: -1, want: math.Inf(-1), test: "-Inf minus minus one"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.test, tc.want, got)
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
		test string
	}{
		{a: 1, b: 2, want: 2, test: "one times two"},
		{a: 2, b: 5, want: 10, test: "two times five"},
		{a: -2, b: -2, want: 4, test: "minus two times minus two"},
		{a: 0, b: -2, want: 0, test: "zero times minus two"},
		{a: -2.6, b: 0.4, want: -1.04, test: "minus two point six minus zero point four"},
		{a: math.Inf(1), b: -1, want: math.Inf(-1), test: "+Inf times minus one"},
		{a: math.Inf(-1), b: -1, want: math.Inf(1), test: "-Inf times minus one"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.test, tc.want, got)
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
		test     string
		errExpected bool
	}{
		{a: 2, b: 1, want: 2, errExpected: false, test: "two divided by 1"},
		{a: 10, b: 5, want: 2, errExpected: false, test: "ten divided by 5"},
		{a: 10, b: 0, want: 0, errExpected: true, test: "error out on divide by zero"},
		{a: 10, b: 0, want: 1, errExpected: true, test: "test error case by failing want vs get"},
		{a: 10, b: 0, want: 1_000_000, errExpected: true, test: "dodgy want to test err"},
		{a: 10, b: 0, want: 999, errExpected: true, test: "dodgy want to test err"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errorReceived := err != nil
		if errorReceived != tc.errExpected {
			t.Fatalf("test %s: expected err %t, got %t", tc.test, tc.errExpected, errorReceived)
		}
		if !errorReceived && tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.test, tc.want, got)
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
		test     string
		errExpected bool
	}{
		{a: 16, want: 4, errExpected: false, test: "square root of 16"},
		{a: -1, want: 0, errExpected: true, test: "error"},
		{a: -1, want: 0, errExpected: true, test: "error out on negative input"},
		{a: -1, want: 1, errExpected: true, test: "test error case by failing want vs get"},
		{a: math.Inf(0), want: math.Inf(0), errExpected: false, test: "test +Inf"},
	}
	t.Parallel()
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errorReceived := err != nil
		if errorReceived != tc.errExpected {
			t.Fatalf("test %s: expected err %t, got %t", tc.test, tc.errExpected, errorReceived)
		}
		if !errorReceived && tc.want != got {
			t.Errorf("test %s: want %f, got %f", tc.test, tc.want, got)
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
