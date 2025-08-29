FROM golang:1.25.0-alpine as builder

LABEL Maintainer="sampozki"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY utils/*.go ./utils/

RUN CGO_ENABLED=0 GOOS=linux go build -o /bottov2

FROM alpine:3.22.1
COPY --from=builder /bottov2 .

CMD ["/app/bottov2"]