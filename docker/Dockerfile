# Build Stage
FROM golang:1.14.3-alpine3.11 AS builder

RUN apk add git

WORKDIR /go/src/gRPCgo
ADD ./ .

#RUN go get -u google.golang.org/grpc
RUN CGO_ENABLED=0 go build -o server.exe

# Deploy Stage
FROM scratch

COPY --from=builder /go/src/gRPCgo/server.exe .

EXPOSE 80
CMD ["./server.exe",  "serve"]