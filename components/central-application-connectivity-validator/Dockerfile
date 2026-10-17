FROM golang:1.22.2-alpine3.18@sha256:d995eb689a0c123590a3d34c65f57f3a118bda3fa26f92da5e089ae7d8fd81a0 as builder

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
