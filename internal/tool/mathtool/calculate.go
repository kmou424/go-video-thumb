package mathtool

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/gookit/goutil/mathutil"
)

func CalculateFloat64(expression string) (float64, error) {
	evaluate, err := calculateExpression(expression)
	if err != nil {
		return 0, err
	}
	res, err := mathutil.ToFloat(evaluate)
	if err != nil {
		return 0, errors.New("convert result of expression to float failed")
	}

	return res, nil
}

func CalculateInt(expression string) (int, error) {
	evaluate, err := calculateExpression(expression)
	if err != nil {
		return 0, err
	}
	res, err := mathutil.ToInt(evaluate)
	if err != nil {
		return 0, errors.New("convert result of expression to int failed")
	}

	return res, nil
}

func calculateExpression(expression string) (any, error) {
	exp, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return nil, errors.New("invalid expression")
	}
	evaluate, err := exp.Evaluate(nil)
	if err != nil {
		return nil, errors.New("calculate expression failed")
	}

	return evaluate, nil
}
