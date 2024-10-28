# syntax=docker/dockerfile:1
FROM golang:1.22 AS builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags osusergo,netgo \
    -trimpath -ldflags='-extldflags=-static -w -s' \
    -o /spacetrouble ./main.go

FROM scratch

# you also need to copy in TLS certificates and timezone data in order to have everything work properly

COPY --from=builder /spacetrouble /spacetrouble

CMD ["/spacetrouble"]
