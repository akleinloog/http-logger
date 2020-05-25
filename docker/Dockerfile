# Build Stage
FROM golang:1.14.3-alpine3.11 AS builder
RUN apk add git
WORKDIR /go/src/http-logger

ADD ./go.mod .
ADD ./go.sum .
RUN go mod download

ADD ./ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httplog

# Deploy Stage
FROM scratch
ARG VERSION="unspecified"
ARG VCS_REF=unspecified
LABEL maintainer="arnoud@kleinloog.ch"
LABEL description="Simple HTTP Logger"
LABEL version=${VERSION}
LABEL vcs_url="https://github.com/akleinloog/http-logger"
LABEL vcs-ref=${VCS_REF}
COPY --from=builder /go/src/http-logger/httplog .
EXPOSE 80
CMD ["./httplog.exe",  "serve"]