FROM golang:1.24.0-alpine3.21

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1

CMD ["oapi-codegen"]
