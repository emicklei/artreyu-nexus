FROM golang:1.4.2-wheezy

RUN go get github.com/emicklei/artreyu

RUN mkdir -p /go/src/github.com/emicklei/artreyu-nexus
WORKDIR /go/src/github.com/emicklei/artreyu-nexus
ADD . /go/src/github.com/emicklei/artreyu-nexus

CMD make docker-build