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
	res := *a.Add(&a, &b)
	op.Result = res
	return res, s.repo.AddOperation(op)
}

func (s operationService) Subtract(a, b big.Float) (big.Float, error) {
	op := Operation{
		OperandA: a,
		OperandB: b,
		Operator: Addition,
	}
	res := *a.Sub(&a, &b)
	op.Result = res
	return res, s.repo.AddOperation(op)
}

func (s operationService) Multiply(a, b big.Float) (big.Float, error) {
	op := Operation{
		OperandA: a,
		OperandB: b,
		Operator: Addition,
	}
	res := *a.Mul(&a, &b)
	op.Result = res
	return res, s.repo.AddOperation(op)
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
	res := *a.Quo(&a, &b)
	op.Result = res
	return res, s.repo.AddOperation(op)
}

func (s operationService) GetRecentN(n int) ([]Operation, error) {
	return s.repo.GetRecentN(n)
}
