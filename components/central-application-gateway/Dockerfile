FROM golang:1.22.3-alpine3.18@sha256:d1a601b64de09e2fa38c95e55838961811d5ca11062a8f4230a5c434b3ae2a34 as builder

ARG DOCK_PKG_DIR=/go/src/github.com/kyma-project/kyma/components/central-application-gateway
WORKDIR $DOCK_PKG_DIR

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $DOCK_PKG_DIR

RUN echo "nobody:x:65534:5534:nobody:/:" > /etc_passwd

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o applicationgateway ./cmd/applicationgateway


FROM eu.gcr.io/kyma-project/external/alpine:3.17.3 as certs
RUN apk add -U --no-cache ca-certificates

FROM scratch
LABEL source=git@github.com:kyma-project/kyma.git

WORKDIR /app

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/kyma-project/kyma/components/central-application-gateway/applicationgateway .

COPY --from=builder /etc_passwd /etc/passwd
USER nobody

CMD ["/app/applicationgateway"]
