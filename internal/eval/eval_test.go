package eval

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr     interface{}
		env      map[string]interface{}
		expected interface{}
	}{
		{expr: 42, env: nil, expected: 42},
		{expr: "x", env: map[string]interface{}{"x": 10}, expected: 10},
		{expr: []interface{}{"+", 1, 2}, env: nil, expected: 3},
		{expr: []interface{}{"-", 5, 3}, env: nil, expected: 2},
		{expr: []interface{}{"*", 2, 3}, env: nil, expected: 6},
		{expr: []interface{}{"/", 6, 2}, env: nil, expected: 3},
		{expr: []interface{}{"<", 1, 2}, env: nil, expected: true},
		{expr: []interface{}{">", 2, 1}, env: nil, expected: true},
		{expr: []interface{}{"=", 2, 2}, env: nil, expected: true},
		{expr: []interface{}{"if", true, 1, 0}, env: nil, expected: 1},
		{expr: []interface{}{"if", false, 1, 0}, env: nil, expected: 0},
		{expr: []interface{}{"define", "y", 20}, env: map[string]interface{}{}, expected: nil},
		{expr: []interface{}{"car", []interface{}{1, 2, 3}}, env: nil, expected: 1},
		{expr: []interface{}{"cdr", []interface{}{1, 2, 3}}, env: nil, expected: []interface{}{2, 3}},
		{expr: []interface{}{"cons", 1, []interface{}{2, 3}}, env: nil, expected: []interface{}{1, 2, 3}},
	}

	for _, test := range tests {
		result := eval(test.expr, test.env)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf(
				"FAIL: eval(%v, %v) = %v (type %T); want %v (type %T)",
				test.expr, test.env, result, result, test.expected, test.expected,
			)
		} else {
			fmt.Printf("PASS: eval(%v) -> %v\n", test.expr, result)
		}
	}
}
