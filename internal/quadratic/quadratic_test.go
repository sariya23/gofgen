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
		coefs       map[string]float64
		undefVar    string
		expectedEq  *quadratic.QuadraticEquation
		expectedErr error
	}{
		{
			caseName: "success create",
			coefs:    map[string]float64{"a": 2.0, "b": 3.0, "c": 4.0},
			undefVar: "x",
			expectedEq: &quadratic.QuadraticEquation{
				EquationCoefficient: map[string]decimal.Decimal{
					"a": decimal.NewFromFloat(2.0),
					"b": decimal.NewFromFloat(3.0),
					"c": decimal.NewFromFloat(4.0),
				},
				UndefinedVariable: "x",
			},
			expectedErr: nil,
		},
		{
			caseName: "success create with point",
			coefs:    map[string]float64{"a": 2.2, "b": -3.1, "c": 4.9},
			undefVar: "x",
			expectedEq: &quadratic.QuadraticEquation{
				EquationCoefficient: map[string]decimal.Decimal{
					"a": decimal.NewFromFloat(2.2),
					"b": decimal.NewFromFloat(-3.1),
					"c": decimal.NewFromFloat(4.9),
				},
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

func TestCreateNewQuadraticNegative(t *testing.T) {
	testCases := []struct {
		caseName    string
		coefs       map[string]float64
		undefVar    string
		expectedEq  *quadratic.QuadraticEquation
		expectedErr error
	}{
		{
			caseName:    "coef by x^2 is zero. Error",
			coefs:       map[string]float64{"a": 0, "b": -3.1, "c": 4.9},
			undefVar:    "x",
			expectedEq:  nil,
			expectedErr: fmt.Errorf("coefficient at the leading term is zero"),
		},
		{
			caseName:    "To much coefs",
			coefs:       map[string]float64{"a": 0, "b": -3.1, "c": 4.9, "d": 3.3},
			undefVar:    "x",
			expectedEq:  nil,
			expectedErr: fmt.Errorf("need 3 coefficients, got %d", 4),
		},
		{
			caseName:    "No enough coefs",
			coefs:       map[string]float64{"a": 0},
			undefVar:    "x",
			expectedEq:  nil,
			expectedErr: fmt.Errorf("need 3 coefficients, got %d", 1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			res, err := quadratic.New(tc.coefs, tc.undefVar)
			if res != nil {
				t.Fatalf("unexpected creating equation")
			}
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("got unexpected error. Expect: (%v), got: (%v)", tc.expectedErr, err)
			}
		})
	}
}
