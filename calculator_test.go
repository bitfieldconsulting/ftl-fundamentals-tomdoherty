package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	comment     string
	err         error
	errExpected bool
}

var testCasesAdd = []testCase {
	{a: 1, b: 2, want: 3, comment: "one add two"},
	{a: 2, b: 5, want: 7, comment: "two add five"},
}

var testCasesSubtract = []testCase{
	{a: 1, b: 2, want: -1, comment: "one minus two"},
	{a: 2, b: 5, want: -3, comment: "two minus five"},
}

var testCasesMultiply = []testCase{
	{a: 1, b: 2, want: 2, comment: "one times two"},
	{a: 2, b: 5, want: 10, comment: "two times five"},
}

var testCasesDivide = []testCase{
	{a: 2, b: 1, want: 2, errExpected: false, comment: "two divided by 1"},
	{a: 10, b: 5, want: 2, errExpected: false, comment: "ten divided by 5"},
	{a: 10, b: 0, want: 0, errExpected: true, comment: "error out on divide by zero"},
}

var testCasesSqrt = []testCase{
	{a: 16, want: 4, errExpected: false, comment: "square root of 16"},
}

func TestAdd(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesAdd {
		want := tc.want
		got := calculator.Add(tc.a, tc.b)
		if want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Errorf("random test: want %f, got %f", want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesSubtract {
		want := tc.want
		got := calculator.Subtract(tc.a, tc.b)
		if want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, want, got)
		}
	}
}

func TestSubtractRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a - b
		got := calculator.Subtract(a, b)
		if want != got {
			t.Errorf("random test: want %f, got %f", want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesMultiply {
		want := tc.want
		got := calculator.Multiply(tc.a, tc.b)
		if want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, want, got)
		}
	}
}

func TestMultiplyRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a * b
		got := calculator.Multiply(a, b)
		if want != got {
			t.Errorf("random test: want %f, got %f", want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesDivide {
		want := tc.want
		got, err := calculator.Divide(tc.a, tc.b)
		errorReceived := err != nil
		if errorReceived != tc.errExpected {
			t.Fatalf("test %s: expected err %t, got %t", tc.comment, tc.errExpected, errorReceived)
		}
		if want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, want, got)
		}
	}
}

func TestDivideRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a / b
		got, _ := calculator.Divide(a, b)
		if want != got {
			t.Errorf("random test: want %f, got %f", want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesSqrt {
		want := tc.want
		got, err := calculator.Sqrt(tc.a)
		errorReceived := err != nil
		if errorReceived != tc.errExpected {
			t.Fatalf("test %s: expected err %t, got %t", tc.comment, tc.errExpected, errorReceived)
		}
		if want != got {
			t.Errorf("test %s: want %f, got %f", tc.comment, want, got)
		}
	}
}

func TestSqrtRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		want := math.Sqrt(a)
		got, _ := calculator.Sqrt(a)
		if want != got {
			t.Errorf("random test: want %f, got %f", want, got)
		}
	}
}
