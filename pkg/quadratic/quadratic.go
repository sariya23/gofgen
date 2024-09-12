package quadratic

import (
	"fmt"
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
// возвращает их в порядке возрастания.
//
// Если дискриминант отрицательный - корней нет (nil);
//
// если дискриминант нулевой - один корень;
//
// если дискриминант положительный - два корня.
// func (qe QuadraticEquation) Solve() []float64 {
// 	var roots []float64
// 	D := math.Pow(float64(qe.B), 2) - 4*qe.A*qe.C
// 	if D == 0 {
// 		root := (-qe.B + math.Sqrt(float64(D))) / (2 * qe.A)
// 		roots = append(roots, root)
// 	} else if D > 0 {
// 		root1 := (-qe.B + math.Sqrt(float64(D))) / (2 * qe.A)
// 		root2 := (-qe.B - math.Sqrt(float64(D))) / (2 * qe.A)
// 		roots = append(roots, root1, root2)
// 		slices.Sort(roots)
// 	} else {
// 		roots = nil
// 	}
// 	return roots
// }
