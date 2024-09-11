package quadratic

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type QuadraticEquation struct {
	A, B, C           decimal.Decimal
	UndefinedVariable string
}

func New(coefficient [3]float64, undefinedVariable string) (*QuadraticEquation, error) {
	coeff := [...]decimal.Decimal{decimal.NewFromFloat(0), decimal.NewFromFloat(0), decimal.NewFromFloat(0)}

	for i := 0; i < 3; i++ {
		d := decimal.NewFromFloat(coefficient[i])
		coeff[i] = d
	}

	if isLeadTermIsZero(coeff) {
		return nil, fmt.Errorf("coefficient at the leading term is zero")
	}

	return &QuadraticEquation{A: coeff[0], B: coeff[1], C: coeff[2], UndefinedVariable: undefinedVariable}, nil
}

func isLeadTermIsZero(coefficient [3]decimal.Decimal) bool {
	return coefficient[0].Equal(decimal.NewFromFloat(0))
}

func (qe QuadraticEquation) String() string {
	stringEquation := fmt.Sprintf("%v%v²", qe.A.String(), qe.UndefinedVariable)

	if qe.B.Compare(decimal.NewFromFloat(0.0)) == -1 {
		stringEquation += fmt.Sprintf(" - %v%v", qe.B.String(), qe.UndefinedVariable)
	} else if qe.B.Compare(decimal.NewFromFloat(0.0)) == 1 {
		stringEquation += fmt.Sprintf(" + %v%v", qe.B.String(), qe.UndefinedVariable)
	}

	if qe.C.Compare(decimal.NewFromFloat(0.0)) == -1 {
		stringEquation += fmt.Sprintf(" - %v", qe.C.String())
	} else if qe.C.Compare(decimal.NewFromFloat(0.0)) == 1 {
		stringEquation += fmt.Sprintf(" + %v", qe.C.String())
	}

	return fmt.Sprintf("%v = 0", stringEquation)
}
