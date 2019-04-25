package lambda

import (
	"encoding/json"
	"testing"

	bindings "github.com/andream16/aws-sdk-go-bindings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
)

func TestUnmarshalDynamoEvent(t *testing.T) {

	t.Run("should return an error because event is empty", func(t *testing.T) {

		err := UnmarshalDynamoEvent(events.DynamoDBEventRecord{}, nil)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because out is not a pointer", func(t *testing.T) {

		err := UnmarshalDynamoEvent(events.DynamoDBEventRecord{
			AWSRegion: "someRegion",
		}, nil)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should successfully unmarshal an event into a given target", func(t *testing.T) {

		type target struct {
			SomeParam string `json:"some_param"`
		}

		var (
			attrv events.DynamoDBAttributeValue
			b     = []byte(`{ 
				"M": {
				  "some_param" : {
					  "S" : "some_val"
				  }
			   }
		  }`)
		)

		err := json.Unmarshal(b, &attrv)
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

		m := attrv.Map()
		event := events.DynamoDBEventRecord{
			Change: events.DynamoDBStreamRecord{
				NewImage: m,
			},
		}

		var trg target

		err = UnmarshalDynamoEvent(event, &trg)
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}
		if "some_val" != trg.SomeParam {
			t.Fatalf("expected value %s, got %s", "some_val", trg.SomeParam)
		}

	})

}
