FROM golang:1.26-alpine AS builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/coffie-api ./cmd/server/main.go

FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
	&& addgroup -S appgroup \
	&& adduser -S -G appgroup -h /home/appuser appuser

WORKDIR /app

COPY --from=builder /out/coffie-api /app/coffie-api

USER appuser

EXPOSE 8080

CMD ["/app/coffie-api"]
