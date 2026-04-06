FROM golang:1.26-alpine

RUN apk add --no-cache git

# Install air for hot-reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod ./
RUN if [ -f go.sum ]; then go mod download; fi

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
