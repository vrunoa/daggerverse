FROM python:3.9.22-alpine3.21

RUN apk update
RUN apk add uv
RUN apk add --no-cache python3
RUN apk add bash

CMD ["uv", "--version"]
