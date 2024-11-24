# Etapa 1: Build
FROM golang:1.22-alpine AS builder

# Instalar dependências básicas
RUN apk add --no-cache git

# Definir diretório de trabalho
WORKDIR /large-app

# Copiar o código fonte e arquivos do módulo Go
COPY . /large-app

# Configurar o módulo Go e compilar o binário
RUN go build -o reserva-salas

# Etapa 2: Runtime
FROM alpine:3.18

# Instalar dependências mínimas para execução
RUN apk add --no-cache ca-certificates

# Definir diretório de trabalho
WORKDIR /small-app

# Copiar o binário da etapa de build
COPY --from=builder /large-app /small-app

# Expor a porta usada pela aplicação
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./reserva-salas"]
