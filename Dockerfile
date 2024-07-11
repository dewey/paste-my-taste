FROM golang:1.20-alpine as builder

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev \
    git \
    bash

ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOGC=off

# Add our code
ADD ./ $GOPATH/src/github.com/dewey/paste-my-taste

# build
WORKDIR $GOPATH/src/github.com/dewey/paste-my-taste
RUN cd $GOPATH/src/github.com/dewey/paste-my-taste && go build -mod=vendor -v -o /paste-my-taste .

# multistage
FROM alpine:latest

# https://stackoverflow.com/questions/33353532/does-alpine-linux-handle-certs-differently-than-busybox#33353762
RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /paste-my-taste /usr/bin/paste-my-taste
COPY --from=builder /web/dist /web/dist
COPY --from=builder /nginx.conf.sigil /nginx.conf.sigil

# Run the image as a non-root user
RUN adduser -D pmt
RUN chmod 0755 /usr/bin/paste-my-taste

USER pmt

# Run the app. CMD is required to run on Heroku
CMD paste-my-taste 