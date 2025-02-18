package calculator

import (
	"github.com/Knetic/govaluate"
)

func Calculate(input string) (any, error) {

	expression, err := govaluate.NewEvaluableExpression(input)
	if err != nil {
		return nil, err
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return nil, err
	}

	return result, nil

}