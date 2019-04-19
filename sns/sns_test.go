package sns

import (
	"testing"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/pkg/errors"
)

func TestNew(t *testing.T) {

	t.Run("should return an error because region is missing", func(t *testing.T) {
		_, err := New("")
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected missing required parameter region error, got %s", err)
		}
	})

}
