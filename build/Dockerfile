FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o app ./cmd/app

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/app /app/app

EXPOSE 8080
ENTRYPOINT ["/app/app"]
