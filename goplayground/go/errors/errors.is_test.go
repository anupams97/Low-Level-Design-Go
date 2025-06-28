package errors

import (
	"errors"
	"fmt"
	"testing"
)

var ErrUnsupported = errors.New("this mode is unsupported")

func operation(mode int) (string, error) {
	if mode < 0 {
		return "", fmt.Errorf("err :%w", ErrUnsupported)
	}
	return fmt.Sprintf("running operation with mode :%d", mode), nil
}

func TestOperationUnsupported(t *testing.T) {
	_, err := operation(-1)
	if err == nil {
		t.Fatal("expected")
	}
	if !errors.Is(err, ErrUnsupported) {
		t.Fatalf("expected unsupported error, got %v", err)
	}
}
