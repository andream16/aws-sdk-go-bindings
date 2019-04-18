package sns

import "testing"

func TestNew(t *testing.T) {

	t.Run("should return an error because region is missing", func(t *testing.T) {
		_, err := New("")
		if err == nil {
			t.Fatal("expected missing required parameter region error, got nil")
		}
	})

}
