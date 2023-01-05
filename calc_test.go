package deka

import (
	"testing"

	"github.com/spatecon/deka/ast"
	"github.com/spatecon/deka/lexer"
	"github.com/spatecon/deka/parser"
)

type TI struct {
	src    string
	expect int64
}

var testData = []*TI{
	{"1 + 1", 2},
	{"1 * 1", 1},
	{"1 + 2 * 3", 7},
	{"100 + (55 * 5) + 2 * 3", 381},
}

func Test1(t *testing.T) {
	p := parser.NewParser()
	for _, ts := range testData {
		s := lexer.NewLexer([]byte(ts.src))
		output, err := p.Parse(s)
		if err != nil {
			t.Error(err)
		}

		value, err := output.(ast.IExpr).Eval()
		if err != nil {
			t.Error(err)
		}

		if value != ts.expect {
			t.Errorf("Error: %s = %d. Expected %d\n", ts.src, value, ts.expect)
		}
	}
}
