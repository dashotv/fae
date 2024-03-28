package fae

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := New("test")
	assert.Error(t, err)
	assert.Equal(t, "test", err.Error())
}

func TestErrorf(t *testing.T) {
	err := Errorf("test %s", "error")
	assert.Error(t, err)
	assert.Equal(t, "test error", err.Error())
}

func TestWrap(t *testing.T) {
	err := New("test")
	err = Wrap(err, "wrapped")
	assert.Error(t, err)
	assert.Equal(t, "wrapped: test", err.Error())
}

func TestWrapf(t *testing.T) {
	err := New("test")
	err = Wrapf(err, "wrapped %s", "error")
	assert.Error(t, err)
	assert.Equal(t, "wrapped error: test", err.Error())
}

func TestStackTrace(t *testing.T) {
	err := New("test")
	stack := StackTrace(err)
	assert.NotEmpty(t, stack)
}

func openError() error {
	_, err := os.ReadFile("fileDoesNotExist")
	return Wrap(err, "opening file")
}

func TestCause(t *testing.T) {
	err := openError()
	assert.Error(t, err)

	orig := Cause(err)
	assert.Error(t, orig)
	assert.NotEqual(t, err, orig)
	assert.Equal(t, "no such file or directory", orig.Error())
}

func TestStackTrace2(t *testing.T) {
	err := openError()
	stack := StackTrace(Wrap(err, "test"))
	assert.NotEmpty(t, stack)
	for _, f := range stack {
		fmt.Printf("  %s\n", f)
	}
}

func TestErrorStack(t *testing.T) {
	err := openError()
	assert.Error(t, err)
	stack := ErrorStack(Wrap(err, "test"))
	assert.NotEmpty(t, stack)
	fmt.Println(stack)
}
