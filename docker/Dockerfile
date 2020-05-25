ARG VERSION=unspecified
ARG GIT_URL=https://github.com/akleinloog/http-logger
ARG GIT_COMMIT=unspecified

# Build Stage
FROM golang:1.14.3-alpine3.11 AS builder
RUN apk add git
WORKDIR /go/src/http-logger
ADD ./ .
RUN CGO_ENABLED=0 go build -o httplog

# Deploy Stage
FROM scratch
LABEL maintainer="arnoud@kleinloog.ch"
LABEL version=$VERSION
LABEL git_url=$GIT_URL
LABEL git_commit=$GIT_COMMIT
COPY --from=builder /go/src/http-logger/httplog .
EXPOSE 80
CMD ["./httplog.exe",  "serve"]