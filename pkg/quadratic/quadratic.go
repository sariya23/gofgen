package quadratic

import (
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"
)

// QuadraticEquation хранит основную информацию о квадратном уравнении.
// Такую как коэффициенты, ответ, название неизвестной переменной.
//
// - A - коэффициент при старшем члене. Не может быть 0;
//
// - B - коэффициент при среднем члене;
//
// - C - свободный член.
//
// - UndefinedVariable - название неизвестной переменной. Например, x.
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
	if coeff[0].IsZero() {
		return nil, fmt.Errorf("coefficient at the leading term is zero")
	}

	return &QuadraticEquation{A: coeff[0], B: coeff[1], C: coeff[2], UndefinedVariable: undefinedVariable}, nil
}

// String возвращает сторокове представление квадратного уравнения.
// Например, 3.4x² + 5x = 4.
func (qe QuadraticEquation) String() string {
	stringEquation := fmt.Sprintf("%v%v²", qe.A.String(), qe.UndefinedVariable)
	if !qe.B.IsZero() {
		stringEquation += fmt.Sprintf("%v%v", parseCoefficient(qe.B), qe.UndefinedVariable)
	}
	if !qe.C.IsZero() {
		stringEquation += parseCoefficient(qe.C)
	}
	return fmt.Sprintf("%v = 0", stringEquation)
}

func parseCoefficient(v decimal.Decimal) string {
	var result string
	if v.IsNegative() {
		result = strings.ReplaceAll(v.String(), "-", " - ")
	} else if v.IsPositive() || v.IsZero() {
		result = fmt.Sprintf(" + %v", v.String())
	}
	return result
}

// Solve решанет квадртаное уравнение. В случае, когда корня два,
// возвращает их в порядке возрастания. Округляет до 2х знаков после запятой.
//
// Если дискриминант отрицательный - корней нет (nil);
//
// если дискриминант нулевой - один корень;
//
// если дискриминант положительный - два корня.
func (qe QuadraticEquation) Solve() []decimal.Decimal {
	var roots []decimal.Decimal
	doubleA := decimal.NewFromInt(2).Mul(qe.A)
	D := qe.B.Pow(decimal.NewFromInt(2)).Sub(decimal.NewFromInt(4).Mul(qe.A).Mul(qe.C))
	if D.IsZero() {
		root := (qe.B.Neg()).Div(doubleA)
		roots = append(roots, root)
	} else if D.IsPositive() {
		f, _ := D.Float64()
		DRoot := decimal.NewFromFloat(math.Sqrt(f))
		root1 := (qe.B.Neg().Add(DRoot)).Div(doubleA)
		root2 := (qe.B.Neg().Sub(DRoot)).Div(doubleA)
		if root1.GreaterThan(root2) {
			roots = append(roots, root2, root1)
		} else {
			roots = append(roots, root1, root2)
		}
	} else {
		roots = nil
	}
	return roots
}
