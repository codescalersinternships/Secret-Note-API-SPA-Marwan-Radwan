FROM golang:1.23 AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy \
    && go build cmd/main.go
FROM scratch AS ginserver
COPY --from=builder /app/ /
EXPOSE 8080
ENTRYPOINT ["/main"]