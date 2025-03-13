FROM golang:1.24@sha256:8678013a2add364dc3d5df2acc2b36893fbbd60ebafa5d5149bc22158512f021

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