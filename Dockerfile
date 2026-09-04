# Build
FROM golang:1.21-alpine AS golang

WORKDIR /app

# Copiar os arquivos Go
COPY http-server-projeto-korp.go .

# Compilar a aplicação
RUN go build -o http-server-projeto-korp http-server-projeto-korp.go

# Executar
FROM alpine:latest

WORKDIR /root/

# Copiar o binário compilado do stage anterior
COPY --from=golang /app/http-server-projeto-korp .

# Expor a porta 8080
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./http-server-projeto-korp"]
