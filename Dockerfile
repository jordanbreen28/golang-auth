
FROM golang:1.19

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go install github.com/pilu/fresh@latest

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

EXPOSE 8080

ENTRYPOINT ["fresh"]
