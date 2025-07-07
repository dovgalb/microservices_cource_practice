FROM golang:1.23.6-alpine AS builder

COPY . /github.com/dovgalb/microservices_cource_practice/week_2/grpc/source/
WORKDIR /github.com/dovgalb/microservices_cource_practice/week_2/grpc/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/dovgalb/microservices_cource_practice/week_2/grpc/source/bin/crud_server .

CMD ["./crud_server"]