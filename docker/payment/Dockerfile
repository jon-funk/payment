FROM golang:1.20.1

COPY . /go/src/github.com/microservices-demo/payment/

WORKDIR /go/src/github.com/microservices-demo/payment/

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app github.com/microservices-demo/payment/cmd/paymentsvc

FROM alpine:3.4

WORKDIR /
COPY --from=0 /app /app

ENV SERVICE_USER=myuser \
    SERVICE_UID=10001 \
    SERVICE_GROUP=mygroup \
    SERVICE_GID=10001

RUN addgroup -g ${SERVICE_GID} ${SERVICE_GROUP} && \
    adduser -g "${SERVICE_NAME} user" -D -H -G ${SERVICE_GROUP} -s /sbin/nologin -u ${SERVICE_UID} ${SERVICE_USER} && \
    chmod +x /app && \
    chown -R ${SERVICE_USER}:${SERVICE_GROUP} /app

USER ${SERVICE_USER}

CMD ["/app", "-port=8080"]

EXPOSE 8080
