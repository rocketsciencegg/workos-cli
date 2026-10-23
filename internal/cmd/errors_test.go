package cmd

import (
	"errors"
	"testing"
)

func TestWrapError(t *testing.T) {
	t.Run("nil error remains nil", func(t *testing.T) {
		if got := wrapError(nil, "operation failed"); got != nil {
			t.Fatalf("wrapError(nil, ...) = %v, want nil", got)
		}
	})

	t.Run("wrapped error preserves message and chain", func(t *testing.T) {
		cause := errors.New("cause")
		got := wrapError(cause, "operation failed")

		if want := "operation failed: cause"; got.Error() != want {
			t.Errorf("wrapped error = %q, want %q", got.Error(), want)
		}
		if !errors.Is(got, cause) {
			t.Errorf("wrapped error %v does not match cause %v", got, cause)
		}
	})
}
