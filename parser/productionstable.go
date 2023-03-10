// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"github.com/spatecon/deka/ast"
	"github.com/spatecon/deka/token"
)

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String:     `S' : Calc	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Calc : Expr	<<  >>`,
		Id:         "Calc",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Expr : Expr "+" Term	<< ast.NewExpr(X[0].(ast.IExpr), X[2].(ast.IExpr)), nil >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewExpr(X[0].(ast.IExpr), X[2].(ast.IExpr)), nil
		},
	},
	ProdTabEntry{
		String:     `Expr : Term	<<  >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Term : Term "*" Factor	<< ast.NewTerm(X[0].(ast.IExpr), X[2].(ast.IExpr)), nil >>`,
		Id:         "Term",
		NTType:     3,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewTerm(X[0].(ast.IExpr), X[2].(ast.IExpr)), nil
		},
	},
	ProdTabEntry{
		String:     `Term : Factor	<<  >>`,
		Id:         "Term",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Factor : "(" Expr ")"	<< X[1], nil >>`,
		Id:         "Factor",
		NTType:     4,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String:     `Factor : int	<< ast.NewLiteralInt(X[0].(*token.Token).Lit) >>`,
		Id:         "Factor",
		NTType:     4,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewLiteralInt(X[0].(*token.Token).Lit)
		},
	},
	ProdTabEntry{
		String:     `Factor : id	<< ast.NewID(X[0].(*token.Token).Lit) >>`,
		Id:         "Factor",
		NTType:     4,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewID(X[0].(*token.Token).Lit)
		},
	},
}
