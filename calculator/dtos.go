package calculator

import (
	"math/big"
)

type OperandsDto struct {
	OperandA big.Float `json:"operandA"`
	OperandB big.Float `json:"operandB"`
}

type ResultDto struct {
	OperandA big.Float
	OperandB big.Float
	Operator Operator
	Result   big.Float
}

func (r ResultDto) MarshalJSON() ([]byte, error) {
	str := r.OperandA.Text('f', -1) + " " + r.Operator.String() + " " + r.OperandB.Text('f', -1) + " = " + r.Result.Text('f', -1)
	return []byte(`"` + str + `"`), nil
}
