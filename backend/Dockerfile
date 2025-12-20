FROM golang:1.24 AS builder

WORKDIR /backend

COPY go.mod go.sum /backend/

RUN go mod tidy

COPY . /backend

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /backend/server ./cmd/

FROM alpine

EXPOSE 8080

COPY --from=builder /backend/server /server

ENTRYPOINT ["/server"]
