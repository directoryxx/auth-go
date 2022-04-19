package helper

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicIfError(t *testing.T) {
	err := errors.New("Test")
	assert.Panics(t, func() { PanicIfError(err) }, "The code did not panic")
}
