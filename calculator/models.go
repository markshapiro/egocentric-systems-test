package calculator

import (
	"math/big"
)

type Operator int

const (
	Addition Operator = iota
	Subtraction
	Multiplication
	Division
)

func (d Operator) String() string {
	return [...]string{"+", "-", "*", "/"}[d]
}

type Operation struct {
	OperandA big.Float
	OperandB big.Float
	Operator Operator
	Result   big.Float
}
