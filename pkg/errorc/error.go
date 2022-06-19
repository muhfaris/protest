package e

import (
	"errors"
	"fmt"
)

type E string

var (
	ErrReadTemplate E = "error read template files"
	ErrNotNumber    E = "error param is not number type"

	// random error
	ErrUnknownRandom E = "error unknown random function"
)

func (e E) Errorf(err error) error {
	return fmt.Errorf("%s (%s)", e, err)
}

func (e E) Error() error {
	return errors.New(string(e))
}
