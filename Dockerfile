FROM golang:1.11-alpine

ENV Jenkins=/go/src/github.com/Jenkins

COPY . $Jenkins/

RUN cd /go/src/github.com/Jenkins \
&& go build

FROM alpine

ENV Jenkins=/go/src/github.com/Jenkins
COPY --from=0  $Jenkins/Jenkins /usr/bin
WORKDIR /data

CMD ["Jenkins"]