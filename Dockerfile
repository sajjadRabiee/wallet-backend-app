FROM repo-afra.snappfood.dev/golang:1.19-alpine as builder
WORKDIR /app/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./e-wallet-api main.go

FROM repo-afra.snappfood.dev/alpine:latest
WORKDIR /app
COPY --from=builder /app/e-wallet-api ./
COPY --from=builder /app/.env ./
COPY --from=builder /app/pkg/swaggerui ./pkg/swaggerui

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./e-wallet-api"]