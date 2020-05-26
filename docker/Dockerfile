# Build Stage
FROM golang:1.14.3-alpine3.11 AS builder
RUN apk add git
WORKDIR /go/src/http-logger
ADD ./go.* ./
RUN go mod download
ADD ./ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httplog

# Deploy Stage
FROM scratch
ARG VERSION
ARG VCS_REF
LABEL maintainer="arnoud@kleinloog.ch" \
    description="Simple HTTP Logger" \
    version=${VERSION} \
    vcs_url="https://github.com/akleinloog/http-logger" \
    vcs-ref=${VCS_REF}
COPY --from=builder /go/src/http-logger/httplog .
EXPOSE 80
CMD ["./httplog",  "serve"]