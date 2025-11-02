# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /server /server

EXPOSE 8080

CMD ["/server"]
