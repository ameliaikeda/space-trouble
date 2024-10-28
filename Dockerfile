# syntax=docker/dockerfile:1
FROM golang:1.22 AS builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags osusergo,netgo \
    -trimpath -ldflags='-extldflags=-static -w -s' \
    -o /spacetrouble ./cmd/main.go

FROM scratch

COPY --from=builder /spacetrouble /spacetrouble

CMD ["/spacetrouble"]
