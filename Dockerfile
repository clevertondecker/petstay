# Etapa 1: Builder (Imagem com Go)
FROM golang:1.22.1-alpine AS builder

# Define o diretório de trabalho no container
WORKDIR /app

# Copia o go.mod e go.sum para o container (para otimizar o download das dependências)
COPY go.mod go.sum ./

# Baixa as dependências do Go
RUN go mod tidy

# Copia todo o código fonte para o container
COPY . .

# Compila o código Go em um binário chamado 'app'
RUN go build -o app .

# Etapa 2: Imagem Final (Executável Enxuto)
FROM alpine:latest

# Instalar Go no container final
RUN apk --no-cache add wget curl && \
    wget https://golang.org/dl/go1.23.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xvzf go1.23.1.linux-amd64.tar.gz && \
    rm go1.23.1.linux-amd64.tar.gz

# Configurar o PATH para o Go
ENV PATH=$PATH:/usr/local/go/bin

# Instalar o cliente MySQL (caso precise)
RUN apk --no-cache add mysql-client

# Define o diretório de trabalho no container
WORKDIR /app/

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/app .

# Garantir permissões de execução no binário
RUN chmod +x /app

# Expondo a porta 8080 para acessar o app
EXPOSE 8080

# Comando para rodar o binário compilado
CMD ["./app"]
