ARG VERSION=unspecified
ARG VCS_URL=https://github.com/akleinloog/http-logger
ARG VCS_REF=unspecified
ARG HTTP_PORT=80

# Build Stage
FROM golang:1.14.3-alpine3.11 AS builder
RUN apk add git
WORKDIR /go/src/http-logger
ADD ./ .
RUN CGO_ENABLED=0 go build -o httplog

# Deploy Stage
FROM scratch
LABEL maintainer="arnoud@kleinloog.ch"
LABEL description="Simple HTTP Logger"
LABEL version=$VERSION
LABEL vcs_url=$VCS_URL
LABEL vcs-ref=$VCS_REF
COPY --from=builder /go/src/http-logger/httplog .
EXPOSE $HTTP_PORT
CMD ["./httplog.exe",  "serve"]