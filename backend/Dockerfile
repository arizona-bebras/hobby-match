FROM golang as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY *.go .
RUN go build -v -o /usr/local/bin/app ./...

WORKDIR /app

COPY pb_hooks /app/pb_hooks

EXPOSE 8080

CMD ["app", "serve", "--http=0.0.0.0:8080"]
