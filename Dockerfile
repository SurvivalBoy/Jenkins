FROM golang:1.11-alpine

ENV Jenkins=/go/src/jenkins

COPY . $Jenkins/

RUN cd /go/src/jenkins \
&& go build

FROM alpine

ENV Jenkins=/go/src/jenkins
COPY --from=0  $Jenkins /usr/bin
WORKDIR /data

CMD ["jenkins"]
