# Etapa 1: build do Go
FROM golang:1.24.5-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o server .


# Etapa 2: usar headless-shell já pronto
FROM chromedp/headless-shell:stable

WORKDIR /app

# Copia apenas o binário Go
COPY --from=builder /app/server /app/server

EXPOSE 8080
CMD ["./server"]
