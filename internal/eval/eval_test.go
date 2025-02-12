package eval

import (
	"fmt"
	"go_lisp_interpreter/internal/parser"
	"reflect"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr     parser.Node
		env      *Env
		expected interface{}
	}{
		{expr: parser.Node{Type: "ATOM", Value: "42"}, env: NewEnv(nil), expected: 42},
		{expr: parser.Node{Type: "ATOM", Value: "x"}, env: func() *Env { env := NewEnv(nil); env.Set("x", 10); return env }(), expected: 10},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "+"},
			{Type: "ATOM", Value: "1"},
			{Type: "ATOM", Value: "2"},
		}}, env: NewEnv(nil), expected: 3},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "-"},
			{Type: "ATOM", Value: "5"},
			{Type: "ATOM", Value: "3"},
		}}, env: NewEnv(nil), expected: 2},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "*"},
			{Type: "ATOM", Value: "2"},
			{Type: "ATOM", Value: "3"},
		}}, env: NewEnv(nil), expected: 6},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "/"},
			{Type: "ATOM", Value: "6"},
			{Type: "ATOM", Value: "2"},
		}}, env: NewEnv(nil), expected: 3},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "define"},
			{Type: "ATOM", Value: "y"},
			{Type: "ATOM", Value: "20"},
		}}, env: NewEnv(nil), expected: nil},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "if"},
			{Type: "ATOM", Value: "true"},
			{Type: "ATOM", Value: "1"},
			{Type: "ATOM", Value: "0"},
		}}, env: NewEnv(nil), expected: 1},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "if"},
			{Type: "ATOM", Value: "false"},
			{Type: "ATOM", Value: "1"},
			{Type: "ATOM", Value: "0"},
		}}, env: NewEnv(nil), expected: 0},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "car"},
			{Type: "LIST", Children: []parser.Node{
				{Type: "ATOM", Value: "1"},
				{Type: "ATOM", Value: "2"},
				{Type: "ATOM", Value: "3"},
			}},
		}}, env: NewEnv(nil), expected: "1"},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "cdr"},
			{Type: "LIST", Children: []parser.Node{
				{Type: "ATOM", Value: "1"},
				{Type: "ATOM", Value: "2"},
				{Type: "ATOM", Value: "3"},
			}},
		}}, env: NewEnv(nil), expected: []interface{}{"2", "3"}},
		{expr: parser.Node{Type: "LIST", Children: []parser.Node{
			{Type: "ATOM", Value: "cons"},
			{Type: "ATOM", Value: "1"},
			{Type: "LIST", Children: []parser.Node{
				{Type: "ATOM", Value: "2"},
				{Type: "ATOM", Value: "3"},
			}},
		}}, env: NewEnv(nil), expected: []interface{}{"1", "2", "3"}},
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
