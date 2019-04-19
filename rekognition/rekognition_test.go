package rekognition

import (
	"testing"

	bindings "github.com/andream16/aws-sdk-go-bindings"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"

	"github.com/pkg/errors"
)

type mockRekognitionClient struct {
	rekognitioniface.RekognitionAPI
}

func (*mockRekognitionClient) CompareFaces(*rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
	return nil, nil
}

func (*mockRekognitionClient) DetectFaces(*rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error) {
	return nil, nil
}

func (*mockRekognitionClient) DetectText(*rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error) {
	return nil, nil
}

type mockFailingRekognitionClient struct {
	rekognitioniface.RekognitionAPI
}

func (*mockFailingRekognitionClient) CompareFaces(*rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
	return nil, errors.New("some error")
}

func (*mockFailingRekognitionClient) DetectFaces(*rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error) {
	return nil, errors.New("some error")
}

func (*mockFailingRekognitionClient) DetectText(*rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error) {
	return nil, errors.New("some error")
}

func TestRekognition_CompareFaces(t *testing.T) {

	t.Run("should return an error because source is empty", func(t *testing.T) {

		r := &Rekognition{}

		_, err := r.CompareFaces(nil, nil, 0)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because target is empty", func(t *testing.T) {

		r := &Rekognition{}

		_, err := r.CompareFaces([]byte{1, 131, 20}, nil, 0)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because compare faces failed", func(t *testing.T) {

		mockSvc := &mockFailingRekognitionClient{}

		r := &Rekognition{
			rekognition: mockSvc,
		}

		_, err := r.CompareFaces([]byte{1, 131, 20}, []byte{1}, 1.2)
		if err == nil {
			t.Fatal("expected error because compare faces failed, got nil")
		}

	})

	t.Run("should successfully compare two faces", func(t *testing.T) {

		mockSvc := &mockRekognitionClient{}

		r := &Rekognition{
			rekognition: mockSvc,
		}

		_, err := r.CompareFaces([]byte{1, 131, 20}, []byte{1}, 1.2)
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

func TestRekognition_DetectFaces(t *testing.T) {

	t.Run("should return an error because source is empty", func(t *testing.T) {

		r := &Rekognition{}

		_, err := r.DetectFaces(nil)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because detect faces failed", func(t *testing.T) {

		mockSvc := &mockFailingRekognitionClient{}

		r := &Rekognition{
			rekognition: mockSvc,
		}

		_, err := r.DetectFaces([]byte{1, 131, 20})
		if err == nil {
			t.Fatal("expected error because detect faces failed, got nil")
		}

	})

	t.Run("should successfully detect faces", func(t *testing.T) {

		mockSvc := &mockRekognitionClient{}

		r := &Rekognition{
			rekognition: mockSvc,
		}

		_, err := r.DetectFaces([]byte{1, 131, 20})
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}

func TestRekognition_DetectText(t *testing.T) {

	t.Run("should return an error because source is empty", func(t *testing.T) {

		r := &Rekognition{}

		_, err := r.DetectText(nil)
		if bindings.ErrInvalidParameter != errors.Cause(err) {
			t.Fatalf("expected error %s, got %s", bindings.ErrInvalidParameter, err)
		}

	})

	t.Run("should return an error because detect text failed", func(t *testing.T) {

		mockSvc := &mockFailingRekognitionClient{}

		r := &Rekognition{
			rekognition: mockSvc,
		}

		_, err := r.DetectText([]byte{1, 131, 20})
		if err == nil {
			t.Fatal("expected error because detect faces failed, got nil")
		}

	})

	t.Run("should successfully detect text", func(t *testing.T) {

		mockSvc := &mockRekognitionClient{}

		r := &Rekognition{
			rekognition: mockSvc,
		}

		_, err := r.DetectText([]byte{1, 131, 20})
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}

	})

}
