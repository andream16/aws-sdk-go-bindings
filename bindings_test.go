package bindings

import (
	"testing"

	"github.com/andream16/aws-sdk-go-bindings/internal/format"
)

func TestNew(t *testing.T) {

	t.Run("should return an error because region is missing", func(t *testing.T) {
		_, err := New("")
		if err == nil {
			t.Fatal("expected missing required parameter region error, got nil")
		}
	})

	t.Run("should successfully add some option", func(t *testing.T) {

		const (
			edp    = "someEndpoint"
			region = "eu-central-1"
		)

		withEndpoint := func(cfg *Config) {
			cfg.Endpoint = format.StrToPtr(edp)
		}

		sess, err := New(region, withEndpoint)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		if *sess.Config.Region != region {
			t.Fatalf("expected region %s, got %s", region, *sess.Config.Region)
		}

		if *sess.Config.Endpoint != edp {
			t.Fatalf("expected endpoint %s, got %s", edp, *sess.Config.Endpoint)
		}

	})

}
