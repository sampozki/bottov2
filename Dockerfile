FROM golang:1.25.0-alpine

LABEL Maintainer="sampozki"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY utils/*.go ./utils/

RUN CGO_ENABLED=0 GOOS=linux go build -o /bottov2

CMD ["/bottov2"]