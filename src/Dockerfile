FROM golang:1.24-alpine AS builder

WORKDIR /src

# instala git para go modules privados (opcional)
RUN apk add --no-cache git

# copia mod files y descarga deps (cache eficiente)
COPY go.mod go.sum ./
RUN go mod download

# copia el resto del código y compila
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-dragonball ./cmd

##################### 
FROM alpine:3.20

# certificados TLS para llamadas https
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/go-dragonball .

EXPOSE 8080           
ENTRYPOINT ["./go-dragonball"]