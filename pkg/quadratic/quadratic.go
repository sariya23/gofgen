package quadratic

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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
	A, B, C           int
	UndefinedVariable string
}

func New(coefficient [3]int, undefinedVariable string) (*QuadraticEquation, error) {
	if isLeadTermIsZero(coefficient) {
		return nil, fmt.Errorf("coefficient at the leading term is zero")
	}

	return &QuadraticEquation{A: coefficient[0], B: coefficient[1], C: coefficient[2], UndefinedVariable: undefinedVariable}, nil
}

func isLeadTermIsZero(coefficient [3]int) bool {
	return coefficient[0] == 0
}

// String возвращает сторокове представление квадратного уравнения.
// Например, 3.4x² + 5x = 4.
func (qe QuadraticEquation) String() string {
	stringEquation := fmt.Sprintf("%v%v²", strconv.Itoa(qe.A), qe.UndefinedVariable)
	if qe.B != 0 {
		stringEquation += fmt.Sprintf("%v%v", parseCoefficient(qe.B), qe.UndefinedVariable)
	}
	if qe.C != 0 {
		stringEquation += parseCoefficient(qe.C)
	}
	return fmt.Sprintf("%v = 0", stringEquation)
}

func parseCoefficient(v int) string {
	var result string
	if v < 0 {
		result = strings.ReplaceAll(strconv.Itoa(v), "-", " - ")
	} else if v >= 0 {
		result = fmt.Sprintf(" + %v", strconv.Itoa(v))
	}
	return result
}

// Solve решанет квадртаное уравнение.
// Если дискриминант отрицательный - корней нет (nil);
//
// если дискриминант нулевой - один корень;
//
// если дискриминант положительный - два корня.
func (qe QuadraticEquation) Solve() []int {
	var roots []int
	D := int(math.Pow(float64(qe.B), 2)) - 4*qe.A*qe.C
	if D == 0 {
		root := (-qe.B + int(math.Sqrt(float64(D)))) / (2 * qe.A)
		roots = append(roots, root)
	} else {
		root1 := (-qe.B + int(math.Sqrt(float64(D)))) / (2 * qe.A)
		root2 := (-qe.B - int(math.Sqrt(float64(D)))) / (2 * qe.A)
		roots = append(roots, root1, root2)
	}
	return roots
}
