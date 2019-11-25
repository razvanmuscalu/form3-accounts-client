FROM golang:1.13.4

ARG goconvey_port

ENV GOCONVEY_PORT=$goconvey_port
ENV APP_SRC_PATH=${GOPATH}/src/github.com/razvanmuscalu/form3-accounts-client

EXPOSE ${GOCONVEY_PORT}

ADD . /go/src/accountapi-client

WORKDIR /go/src/accountapi-client

RUN go get accountapi-client
RUN go get github.com/smartystreets/goconvey/convey
RUN go get github.com/smartystreets/goconvey
RUN go install

CMD goconvey -host=0.0.0.0 -port=${GOCONVEY_PORT} -workDir=${APP_SRC_PATH} -launchBrowser=false