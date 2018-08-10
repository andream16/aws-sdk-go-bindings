FROM golang:1.10-alpine3.7 as builder
WORKDIR /go/src/github.com/andream16/aws-sdk-go-bindings

RUN apk --no-cache add git
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -vendor-only

COPY . ./
RUN go build

FROM alpine:3.7
COPY --from=builder /go/src/andream16/aws-sdk-go-bindings ./
CMD ["./aws-sdk-go-bindings"]
