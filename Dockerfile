# Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar os arquivos Go
COPY main.go .

# Compilar a aplicação
RUN go build -o server main.go

# Executar
FROM alpine:latest

WORKDIR /root/

# Copiar o binário compilado do stage anterior
COPY --from=builder /app/server .

# Expor a porta 8080
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./server"]
