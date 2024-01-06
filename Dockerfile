FROM golang:1.21 as base
FROM base as dev-turbo
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

#CMD ["air", "-c", ".air.toml"]
CMD ["air", "server", "--port", "80"]