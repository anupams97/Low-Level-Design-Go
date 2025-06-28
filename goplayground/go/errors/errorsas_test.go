package errors

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

type UnsupportedOp struct {
	mode int
}

func (e *UnsupportedOp) Error() string {
	return fmt.Sprintf("operation: mode %d is unsupported", e.mode)
}

func op(mode int) (string, error) {
	if mode < 0 {
		return "", &UnsupportedOp{mode: mode}
	}
	return fmt.Sprintf("running operation for mode: %d", mode), nil
}

func Test_op(t *testing.T) {
	_, err := op(-1)
	if err == nil {
		log.Fatal("expected error for unsupported mode")
	}
	// because *UnsupportedOp is implementing the error interface
	var unsupportedErr *UnsupportedOp
	// passing pointer to variable you are searching for
	if errors.As(err, &unsupportedErr) {
		if unsupportedErr.mode != -1 {
			log.Fatal("oh no!")
		}
	}

}
