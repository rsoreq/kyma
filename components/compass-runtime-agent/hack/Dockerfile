FROM golang:1.14@sha256:1a7173b5b9a3af3e29a5837e0b2027e1c438fd1b83bbee8f221355087ad416d6

ENV GO111MODULE=off

WORKDIR /go/src/k8s.io

RUN git clone https://github.com/kubernetes/code-generator.git
# Versions needs to be compatible with Kuberentes client used
# hadolint ignore=DL3003
RUN cd code-generator && git fetch && git checkout release-1.21

RUN go get k8s.io/apimachinery; exit 0

ARG repo="${GOPATH}/src/${PROJECT_MODULE}"

RUN mkdir -p $repo

WORKDIR $GOPATH/src/k8s.io/code-generator

ENV GO111MODULE=on

RUN go mod vendor

ENV GO111MODULE=off

VOLUME $repo