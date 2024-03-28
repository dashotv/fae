package fae

import (
	"fmt"

	"github.com/go-errors/errors"
)

func New(message interface{}) error {
	return errors.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Wrap(err error, message string) error {
	return errors.WrapPrefix(err, message, 1)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.WrapPrefix(err, fmt.Sprintf(format, args...), 1)
}

func Cause(err error) error {
	type unwrapper interface {
		Unwrap() error
	}

	for err != nil {
		cause, ok := err.(unwrapper)
		if !ok {
			break
		}
		err = cause.Unwrap()
	}
	return err
}

func StackTrace(err error) []string {
	if err == nil {
		return nil
	}

	list := errors.Wrap(err, 1).StackFrames()
	cause := Cause(err)
	stack := []string{cause.Error()}

	for _, frame := range list {
		// stack = append(stack, fmt.Sprintf("%s.%s %s:%d (0x%x)", frame.Package, frame.Name, frame.File, frame.LineNumber, frame.ProgramCounter))
		stack = append(stack, fmt.Sprintf("%s.%s %s:%d", frame.Package, frame.Name, frame.File, frame.LineNumber))
	}
	return stack
}

func ErrorStack(err error) string {
	return string(errors.Wrap(err, 1).Stack())
}
