package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwarp(t *testing.T) {
	err := fmt.Errorf("some error: %w", errors.New("wrapped error"))
	t.Log(err)
	underlyingError := errors.Unwrap(err)
	t.Log(underlyingError)
}
