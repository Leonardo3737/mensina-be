# Use uma imagem oficial do Go
FROM golang:1.23-alpine3.19

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie todos os arquivos do projeto para o container
COPY . .

# Baixe as dependências
RUN go mod download

# Compile a aplicação
RUN go build -o main .

# Exponha a porta em que o servidor vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]
