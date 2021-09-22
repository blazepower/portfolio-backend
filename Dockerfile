# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY /Handlers/go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8080

ENTRYPOINT ["/app/portfolio-backend"]