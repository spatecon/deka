package ast

import "github.com/spatecon/deka/util"

type IExpr interface {
	Eval() (int64, error)
}

type Expr struct {
	Left  IExpr
	Right IExpr
}

func NewExpr(left IExpr, right IExpr) *Expr {
	return &Expr{Left: left, Right: right}
}

func (e *Expr) Eval() (int64, error) {
	left, err := e.Left.Eval()
	if err != nil {
		return 0, err
	}

	right, err := e.Right.Eval()
	if err != nil {
		return 0, err
	}

	return left + right, nil
}

type Term struct {
	Left  IExpr
	Right IExpr
}

func NewTerm(left IExpr, right IExpr) *Term {
	return &Term{Left: left, Right: right}
}

func (e *Term) Eval() (int64, error) {
	left, err := e.Left.Eval()
	if err != nil {
		return 0, err
	}

	right, err := e.Right.Eval()
	if err != nil {
		return 0, err
	}

	return left * right, nil
}

type ID struct {
	Value string
}

func NewID(value []byte) (*ID, error) {
	// TODO: check if id exists
	return &ID{Value: string(value)}, nil
}

func (i *ID) Int() int64 {
	if i.Value == "visa.score" {
		return 600
	}

	return 0
}

func (i *ID) Eval() (int64, error) {
	return i.Int(), nil
}

type LiteralInt struct {
	Value int64
}

func NewLiteralInt(value []byte) (*LiteralInt, error) {
	v, err := util.IntValue(value)
	if err != nil {
		return nil, err
	}

	return &LiteralInt{Value: v}, nil
}

func (l *LiteralInt) Int() int64 {
	return l.Value
}

func (l *LiteralInt) Eval() (int64, error) {
	return l.Int(), nil
}
