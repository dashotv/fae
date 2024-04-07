package fae

import (
	"fmt"

	"github.com/go-errors/errors"
)

type Error = errors.Error
type Errorish interface {
	Error() string
}

func New(message interface{}) *Error {
	return errors.New(message)
}

func Errorf(format string, args ...interface{}) *Error {
	return errors.Errorf(format, args...)
}

func Wrap(err Errorish, message string) *Error {
	if err == nil {
		return nil
	}
	return errors.WrapPrefix(err, message, 2) // TODO: this might need to be 2, need to test stacktraces more
}

func Wrapf(err Errorish, format string, args ...interface{}) *Error {
	if err == nil {
		return nil
	}
	return errors.WrapPrefix(err, fmt.Sprintf(format, args...), 2) // TODO: this might need to be 2, need to test stacktraces more
}

func Cause(err Errorish) *Error {
	if err == nil {
		return nil
	}

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
	return err.(*Error)
}

func StackTrace(err Errorish) []string {
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
