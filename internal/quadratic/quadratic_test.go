package quadratic_test

import (
	"fmt"
	"gofgen/internal/quadratic"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
)

// TestCreateNewEquationSuccess проверяет, что
// квадратное уравнение создается при передаче
// верных параметров.
func TestCreateNewEquationSuccess(t *testing.T) {
	testCases := []struct {
		caseName    string
		coefs       [3]float64
		undefVar    string
		expectedEq  *quadratic.QuadraticEquation
		expectedErr error
	}{
		{
			caseName: "success create",
			coefs:    [3]float64{2.1, 2.0, 0},
			undefVar: "x",
			expectedEq: &quadratic.QuadraticEquation{
				A:                 decimal.NewFromFloat(2.1),
				B:                 decimal.NewFromFloat(2.0),
				C:                 decimal.NewFromFloat(0),
				UndefinedVariable: "x",
			},
			expectedErr: nil,
		},
		{
			caseName: "success create with point",
			coefs:    [3]float64{123.90, 1.1, -90.2},
			undefVar: "x",
			expectedEq: &quadratic.QuadraticEquation{
				A:                 decimal.NewFromFloat(123.90),
				B:                 decimal.NewFromFloat(1.1),
				C:                 decimal.NewFromFloat(-90.2),
				UndefinedVariable: "x",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			res, err := quadratic.New(tc.coefs, tc.undefVar)
			if err != nil {
				t.Errorf("got unexpected error: (%v)", err)
			}
			if diff := cmp.Diff(*tc.expectedEq, *res); diff != "" {
				t.Fatal(diff)
			}
			if res.UndefinedVariable != tc.undefVar {
				t.Errorf("Expected: (%s), got: (%s)", res.UndefinedVariable, tc.undefVar)
			}
		})
	}
}

// TestCannotCreateQuadraticEquationWithZeroLeadTerm проверяет,
// что квадратное уравнение не создаться, если коэффициент при старшем
// члене 0.
func TestCannotCreateQuadraticEquationWithZeroLeadTerm(t *testing.T) {
	testCase := struct {
		caseName    string
		coefs       [3]float64
		undefVar    string
		expectedEq  *quadratic.QuadraticEquation
		expectedErr error
	}{
		coefs:       [3]float64{0.0, -3.1, 4.9},
		undefVar:    "x",
		expectedErr: fmt.Errorf("coefficient at the leading term is zero"),
	}

	res, err := quadratic.New(testCase.coefs, testCase.undefVar)
	if res != nil {
		t.Fatalf("unexpected creating equation")
	}
	if err.Error() != testCase.expectedErr.Error() {
		t.Errorf("got unexpected error. Expect: (%v), got: (%v)", testCase.expectedErr, err)
	}
}
