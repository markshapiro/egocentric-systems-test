package calculator

import (
	"bufio"
	"os"
	"strings"
)

type OperationRepo interface {
	AddOperation(op Operation) error
	GetRecentN(n int) ([]Operation, error)
}

type operationRepo struct {
	operations []Operation
}

func NewOperationRepo() OperationRepo {
	operations, err := loadFromFile()
	if err != nil {
		panic("failed to load data")
	}

	return &operationRepo{operations}
}

func (repo *operationRepo) AddOperation(op Operation) error {
	repo.operations = append(repo.operations, op)
	return addToFile(op)
}

func (repo operationRepo) GetRecentN(n int) ([]Operation, error) {
	if len(repo.operations) < n {
		n = len(repo.operations)
	}
	return repo.operations[len(repo.operations)-n:], nil
}

func loadFromFile() ([]Operation, error) {
	var operations []Operation
	readFile, err := os.Open("./db.txt")
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		line := fileScanner.Text()

		arr := strings.Split(line, ",")

		var op Operation

		err := op.OperandA.UnmarshalText([]byte(arr[0]))
		if err != nil {
			return nil, err
		}
		err = op.OperandB.UnmarshalText([]byte(arr[1]))
		if err != nil {
			return nil, err
		}
		err = op.Result.UnmarshalText([]byte(arr[3]))
		if err != nil {
			return nil, err
		}

		switch arr[2] {
		case "+":
			op.Operator = Addition
		case "-":
			op.Operator = Subtraction
		case "*":
			op.Operator = Multiplication
		case "/":
			op.Operator = Division
		}

		operations = append(operations, op)
	}
	return operations, nil
}

func addToFile(op Operation) error {
	f, err := os.OpenFile("./db.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	t1, err := op.OperandA.MarshalText()
	if err != nil {
		return err
	}
	t2, err := op.OperandA.MarshalText()
	if err != nil {
		return err
	}
	rs, err := op.OperandA.MarshalText()
	if err != nil {
		return err
	}

	var str = string(t1) + "," + string(t2) + "," + op.Operator.String() + "," + string(rs) + "\n"

	if _, err = f.WriteString(str); err != nil {
		return err
	}
	return nil
}
