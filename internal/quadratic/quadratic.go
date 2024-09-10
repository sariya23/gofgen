package quadratic

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Coefficient = map[string]decimal.Decimal

type QuadraticEquation struct {
	EquationCoefficient Coefficient
	UndefinedVariable   string
}

func New(coefficient map[string]float64, undefinedVariable string) (*QuadraticEquation, error) {
	if coeff := len(coefficient); coeff != 3 {
		return nil, fmt.Errorf("need 3 coefficients, got %d", coeff)
	}
	coeff := make(Coefficient, 3)

	for k, v := range coefficient {
		d := decimal.NewFromFloat(v)
		coeff[k] = d
	}

	return &QuadraticEquation{EquationCoefficient: coeff, UndefinedVariable: undefinedVariable}, nil
}
