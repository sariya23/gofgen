package quadratic_test

import (
	"fmt"
	"gofgen/pkg/quadratic"
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
			caseName: "Success create with positive floats",
			coefs:    [3]float64{2, 2, 0},
			undefVar: "x",
			expectedEq: &quadratic.QuadraticEquation{
				A:                 decimal.NewFromInt(2),
				B:                 decimal.NewFromInt(2),
				C:                 decimal.NewFromInt(0),
				UndefinedVariable: "x",
			},
		},
		{
			caseName: "success create with negative coefficient",
			coefs:    [3]float64{123, 1, -90},
			undefVar: "x",
			expectedEq: &quadratic.QuadraticEquation{
				A:                 decimal.NewFromInt(123),
				B:                 decimal.NewFromInt(1),
				C:                 decimal.NewFromInt(-90),
				UndefinedVariable: "x",
			},
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
		coefs:       [3]float64{0, -3, 4},
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

// TestShowEquationAsString проверяет,
// что уравнение корректно отображается в строковом виде:
//
// - между знаками и значениеями стоит проблел (2x² + 5);
//
// - у страшего члена не должно быть пробела между знаком и значением (-2x²);
//
// - если значение коэффициента n.0, то дробная часть не должна отображаться.
//
// - если значение у коэфициента при среднем члене или у свободного члена 0,
// то он не должен отображться (2x² + 5 = 0);
//
// - если коэффициент при страшем члене положительный, то знак + не отображается (2x²).
func TestShowEquationAsString(t *testing.T) {
	testCases := []struct {
		caseName            string
		equationCoefficient [3]float64
		undefinedVariable   string
		expectedString      string
	}{
		{
			caseName:            "all coefficients are postitive",
			equationCoefficient: [3]float64{2, 23, 5},
			undefinedVariable:   "z",
			expectedString:      "2z² + 23z + 5 = 0",
		},
		{
			caseName:            "all coefficient are negative",
			equationCoefficient: [3]float64{-2, -23, -5},
			undefinedVariable:   "x",
			expectedString:      "-2x² - 23x - 5 = 0",
		},
		{
			caseName:            "zero second coefficient",
			equationCoefficient: [3]float64{2, 0, 5},
			undefinedVariable:   "qwe",
			expectedString:      "2qwe² + 5 = 0",
		},
		{
			caseName:            "zero free term",
			equationCoefficient: [3]float64{3, 5, 0},
			undefinedVariable:   "x",
			expectedString:      "3x² + 5x = 0",
		},
		{
			caseName:            "zero second and free term",
			equationCoefficient: [3]float64{2, 0, 0},
			undefinedVariable:   "x",
			expectedString:      "2x² = 0",
		},
		{
			caseName:            "combination of positive and negative terms",
			equationCoefficient: [3]float64{-2, 10, 2},
			undefinedVariable:   "x",
			expectedString:      "-2x² + 10x + 2 = 0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			eq, err := quadratic.New(tc.equationCoefficient, tc.undefinedVariable)
			if err != nil {
				t.Fatalf("unexpected error: (%v)", err)
			}
			res := eq.String()
			if diff := cmp.Diff(res, tc.expectedString); diff != "" {
				t.Error(diff)
			}
		})
	}
}

// TestSolveQuadraticEquation проверяет,
// что квадратное уравнение решается верно.
// func TestSolveQuadraticEquation(t *testing.T) {
// 	testCases := []struct {
// 		caseName            string
// 		equationCoefficient [3]decimal.Decimal
// 		roots               []decimal.Decimal
// 	}{
// 		{
// 			caseName:            "No roots - nil",
// 			equationCoefficient: [...]decimal.Decimal{2, -1, 1},
// 			roots:               nil,
// 		},
// 		{
// 			caseName:            "Two roots",
// 			equationCoefficient: [...]decimal.Decimal{1, -4, -5},
// 			roots:               []decimal.Decimal{-1, 5},
// 		},
// 		{
// 			caseName:            "One root",
// 			equationCoefficient: [...]decimal.Decimal{1, -2, 1},
// 			roots:               []decimal.Decimal{1},
// 		},
// 		{
// 			caseName:            "Two roots. Decimal",
// 			equationCoefficient: [3]decimal.Decimal{3, 5, 2},
// 			roots:               []decimal.Decimal{5.0 / 6.0, 1.0 / 2.0},
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.caseName, func(t *testing.T) {
// 			eq, err := quadratic.New(tc.equationCoefficient, "x")
// 			if err != nil {
// 				t.Fatalf("unexpected err: (%v)", err)
// 			}
// 			res := eq.Solve()
// 			fmt.Println(res)
// 			if diff := cmp.Diff(res, tc.roots); diff != "" {
// 				t.Error(diff)
// 			}
// 		})
// 	}
// }
