package random

import (
	"math/rand"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
	e "gitlab.sicepat.tech/muhfaris/protest/pkg/errorc"
)

type Expression struct {
	exp *govaluate.EvaluableExpression
}

const randomTag = "random:"

func Random(exp string) (*Expression, error) {
	exp = strings.TrimPrefix(exp, randomTag)

	fns := map[string]govaluate.ExpressionFunction{
		"number": func(arguments ...interface{}) (interface{}, error) {
			var rn int
			if len(arguments) == 2 {
				min, ok := arguments[0].(int)
				if !ok {
					return nil, e.ErrNotNumber.Error()
				}

				max, ok := arguments[1].(int)
				if !ok {
					return nil, e.ErrNotNumber.Error()
				}

				rand.Seed(time.Now().UnixNano())
				rn = rand.Intn(max-min+1) + min
				return rn, nil
			}

			rn = rand.Int()
			return rn, nil
		},
	}

	ve, err := govaluate.NewEvaluableExpressionWithFunctions(exp, fns)
	if err != nil {
		return nil, err
	}

	return &Expression{
		exp: ve,
	}, nil
}
