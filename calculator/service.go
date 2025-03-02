package calculator

import (
	"errors"
	"math/big"
)

var (
	ErrDivisionByZero = errors.New("divizion by zero")
)

type OperationService interface {
	Add(a, b big.Float) (big.Float, error)
	Subtract(a, b big.Float) (big.Float, error)
	Multiply(a, b big.Float) (big.Float, error)
	Divide(a, b big.Float) (big.Float, error)
	GetRecentN(n int) ([]Operation, error)
}

type operationService struct {
	repo OperationRepo
}

func NewOperationService(repo OperationRepo) OperationService {
	return operationService{repo}
}

func (s operationService) Add(a, b big.Float) (big.Float, error) {
	op := Operation{
		OperandA: a,
		OperandB: b,
		Operator: Addition,
	}
	op.Result = *a.Add(&a, &b)
	return op.Result, s.repo.AddOperation(op)
}

func (s operationService) Subtract(a, b big.Float) (big.Float, error) {
	op := Operation{
		OperandA: a,
		OperandB: b,
		Operator: Addition,
	}
	op.Result = *a.Sub(&a, &b)
	return op.Result, s.repo.AddOperation(op)
}

func (s operationService) Multiply(a, b big.Float) (big.Float, error) {
	op := Operation{
		OperandA: a,
		OperandB: b,
		Operator: Addition,
	}
	op.Result = *a.Mul(&a, &b)
	return op.Result, s.repo.AddOperation(op)
}

func (s operationService) Divide(a, b big.Float) (big.Float, error) {
	if b.Cmp(big.NewFloat(0.0)) == 0 {
		return *big.NewFloat(0.0), ErrDivisionByZero
	}
	op := Operation{
		OperandA: a,
		OperandB: b,
		Operator: Addition,
	}
	op.Result = *a.Quo(&a, &b)
	return op.Result, s.repo.AddOperation(op)
}

func (s operationService) GetRecentN(n int) ([]Operation, error) {
	return s.repo.GetRecentN(n)
}
