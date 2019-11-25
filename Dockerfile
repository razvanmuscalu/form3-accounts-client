FROM golang:1.13.4

ADD . /go/src/accountapi-client

WORKDIR /go/src/accountapi-client

RUN go get accountapi-client
RUN go get github.com/smartystreets/goconvey/convey
RUN go install