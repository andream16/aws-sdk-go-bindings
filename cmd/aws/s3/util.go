package s3

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/s3"
)

// ReadImageOutput embeds the result of opening an image and getting its metadata
type ReadImageOutput struct {
	Body        []byte
	ContentType string
	ContentSize int64
}

// SetBody sets ReadImageOutput.Body to the passed body
func (img *ReadImageOutput) SetBody(body []byte) *ReadImageOutput {
	img.Body = body
	return img
}

// SetContentType sets ReadImageOutput.ContentType to the passed contentType
func (img *ReadImageOutput) SetContentType(contentType string) *ReadImageOutput {
	img.ContentType = contentType
	return img
}

// SetContentSize sets ReadImageOutput.ContentSize to the passed contentSize
func (img *ReadImageOutput) SetContentSize(contentSize int64) *ReadImageOutput {
	img.ContentSize = contentSize
	return img
}

// UnmarshalGetObjectOutput extracts bytes from *s3.GetObjectOutput
func UnmarshalGetObjectOutput(input *s3.GetObjectOutput) ([]byte, error) {

	if *input.ContentLength == 0 {
		return nil, errors.New(ErrEmptyContentLength)
	}

	body, bytesErr := ioutil.ReadAll(input.Body)
	if bytesErr != nil {
		return nil, bytesErr
	}
	if len(body) == 0 {
		return nil, errors.New(ErrEmptyBody)
	}

	input.Body = ioutil.NopCloser(bytes.NewReader(body))

	b, err := s3.UnmarshalIOReadCloser(input.Body)
	if err != nil {
		return nil, err
	}

	return b, nil

}

// ReadImage reads an image given its path and returns a *ReadImageOutput containing its body and metadata
func ReadImage(path string) (*ReadImageOutput, error) {

	file, fileErr := os.Open(path)
	if fileErr != nil {
		return nil, fileErr
	}

	defer file.Close()

	fileInfo, fileInfoErr := file.Stat()
	if fileInfoErr != nil {
		return nil, fileInfoErr
	}

	contentSize := fileInfo.Size()
	buffer := make([]byte, contentSize)

	file.Read(buffer)
	contentType := http.DetectContentType(buffer)

	out := new(ReadImageOutput)
	out = out.SetBody(buffer)
	out = out.SetContentType(contentType)
	out = out.SetContentSize(contentSize)

	return out, nil

}
