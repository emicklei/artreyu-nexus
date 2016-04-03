FROM golang:1.6-wheezy

RUN mkdir -p /go/src/github.com/emicklei/artreyu-nexus
WORKDIR /go/src/github.com/emicklei/artreyu-nexus
ADD . /go/src/github.com/emicklei/artreyu-nexus

CMD make build