FROM golang:1.22.3-alpine3.18@sha256:d1a601b64de09e2fa38c95e55838961811d5ca11062a8f4230a5c434b3ae2a34 as builder

ARG DOCK_PKG_DIR=/go/src/github.com/kyma-project/kyma/components/central-application-connectivity-validator
WORKDIR $DOCK_PKG_DIR

COPY . $DOCK_PKG_DIR

RUN echo "nobody:x:65534:65534:nobody:/:" > /etc_passwd

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o centralapplicationconnectivityvalidator ./cmd/centralapplicationconnectivityvalidator


FROM scratch
LABEL source=git@github.com:kyma-project/kyma.git

WORKDIR /app

COPY --from=builder /go/src/github.com/kyma-project/kyma/components/central-application-connectivity-validator/centralapplicationconnectivityvalidator .

COPY --from=builder /etc_passwd /etc/passwd
USER nobody

CMD ["/app/centralapplicationconnectivityvalidator"]
