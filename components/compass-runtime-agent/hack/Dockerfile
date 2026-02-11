FROM golang:1.26@sha256:c83e68f3ebb6943a2904fa66348867d108119890a2c6a2e6f07b38d0eb6c25c5

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