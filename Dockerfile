FROM golang:1.10-alpine as builder
WORKDIR $GOPATH/src/github.com/dewey/paste-my-taste
ADD ./ $GOPATH/src/github.com/dewey/paste-my-taste
RUN apk update && \
    apk upgrade && \
    apk add git
RUN cd $GOPATH/src/github.com/dewey/paste-my-taste && \    
    go build -v -o /paste-my-taste
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /paste-my-taste /paste-my-taste
COPY /web/dist /web/dist
CMD ["/paste-my-taste"]