package s3

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"

	intErr "github.com/andream16/aws-sdk-go-bindings/internal/error"
)

// ReadImageOutput embeds the result of opening an image and getting its metadata
type ReadImageOutput struct {
	// Body is the encoded body of the output
	Body []byte
	// ContentType is the content type of the output
	ContentType string
	// ContentSize is the body size
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
		return nil, intErr.Format(InputContentLength, ErrEmptyContentLength)
	}

	body, bytesErr := ioutil.ReadAll(input.Body)
	if bytesErr != nil {
		return nil, bytesErr
	}
	if len(body) == 0 {
		return nil, intErr.Format(Body, ErrEmptyBody)
	}

	input.Body = ioutil.NopCloser(bytes.NewReader(body))

	b, err := UnmarshalIOReadCloser(input.Body)
	if err != nil {
		return nil, err
	}

	return b, nil

}

// ReadImage reads an image given its path and returns a *ReadImageOutput containing its body and metadata
func ReadImage(path string) (*ReadImageOutput, error) {

	if len(path) == 0 {
		return nil, intErr.Format(Path, ErrEmptyParameter)
	}

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

// NewCreateBucketInput returns a new *s3.CreateBucketInput
func NewCreateBucketInput(bucketName string) (*s3.CreateBucketInput, error) {

	if len(bucketName) == 0 {
		return nil, intErr.Format(BucketName, ErrEmptyParameter)
	}

	out := new(s3.CreateBucketInput)
	out = out.SetBucket(bucketName)

	return out, nil

}

// NewGetObjectInput returns a new *s3.GetObjectInput given a bucket and a source image
func NewGetObjectInput(bucketName, source string) (*s3.GetObjectInput, error) {

	if len(bucketName) == 0 {
		return nil, intErr.Format(BucketName, ErrEmptyParameter)
	}
	if len(source) == 0 {
		return nil, intErr.Format(Source, ErrEmptyParameter)
	}

	out := new(s3.GetObjectInput)
	out = out.SetBucket(bucketName)
	out = out.SetKey(source)

	return out, nil

}

// NewPutObjectInput returns a new *s3.PutObjectInput
func NewPutObjectInput(bucketName, fileName, contentType string, image []byte, size int64) (*s3.PutObjectInput, error) {

	if len(bucketName) == 0 {
		return nil, intErr.Format(BucketName, ErrEmptyParameter)
	}
	if len(fileName) == 0 {
		return nil, intErr.Format(FileName, ErrEmptyParameter)
	}
	if len(contentType) == 0 {
		return nil, intErr.Format(ContentType, ErrEmptyParameter)
	}
	if len(image) == 0 {
		return nil, intErr.Format(Image, ErrEmptyParameter)
	}

	out := new(s3.PutObjectInput)
	out = out.SetBucket(bucketName)
	out = out.SetKey(fileName)
	out = out.SetContentType(contentType)
	out = out.SetBody(bytes.NewReader(image))
	out = out.SetContentLength(size)

	return out, nil

}

// UnmarshalIOReadCloser extracts []byte from input.Body
func UnmarshalIOReadCloser(input io.ReadCloser) ([]byte, error) {

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(input)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
