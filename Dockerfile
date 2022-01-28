# FROM golang:1.16-alpine

# WORKDIR /app

# COPY . . 

# EXPOSE 8080

# CMD [ "go", "run", "./server/server.go" ]

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

COPY . . 

RUN go build -o /GoBackend

EXPOSE 8080

CMD [ "/GoBackend" ]