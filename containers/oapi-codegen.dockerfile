FROM golang:1.24.0-alpine3.21

RUN go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0

CMD ["oapi-codegen"]
