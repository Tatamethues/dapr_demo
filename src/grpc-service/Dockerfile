# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY ["src/grpc-service/go.mod","./"]
COPY ["src/grpc-service/go.sum","./"]
RUN go mod download

COPY ["src/grpc-service/.","."]

RUN go build -o grpc-service

EXPOSE 80

CMD [ "./grpc-service" ]