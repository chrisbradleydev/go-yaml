FROM golang:1.23-alpine
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.[sum] ./
RUN go mod download

COPY . .
RUN go build -o main main.go

CMD ["air", "-c", ".air.toml"]
