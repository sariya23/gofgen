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

func New(coefficient map[string]string, undefinedVariable string) (*QuadraticEquation, error) {
	if coeff := len(coefficient); coeff != 3 {
		return nil, fmt.Errorf("not enough coefficient. Need 3, got %d", coeff)
	}
	coeff := make(Coefficient, 3)

	for k, v := range coefficient {
		d, err := decimal.NewFromString(v)
		if err != nil {
			return nil, fmt.Errorf("cannot parse coef %v: %v", v, err)
		}
		coeff[k] = d
	}

	return &QuadraticEquation{EquationCoefficient: coeff, UndefinedVariable: undefinedVariable}, nil
}
