FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./

EXPOSE 8080
RUN echo $

CMD go run server.go