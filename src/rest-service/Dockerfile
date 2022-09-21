# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY ["src/rest-service/go.mod","./"]
COPY ["src/rest-service/go.sum","./"]
RUN go mod download

COPY ["src/rest-service/.","."]

RUN go build -o rest-service

EXPOSE 80

CMD [ "./rest-service" ]