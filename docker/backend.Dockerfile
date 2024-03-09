FROM golang:1.20-buster as builder

WORKDIR /go/src

COPY ../backend/go.* ./
RUN go mod download

COPY ../backend/ ./
RUN go build -v -o /go/bin/server .

FROM debian:buster-slim

WORKDIR /go/bin

COPY --from=builder /go/bin/server ./server

ENTRYPOINT ["/go/bin/server"]
