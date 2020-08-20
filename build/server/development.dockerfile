FROM golang:1.14 AS builder

ARG SERVICE_NAME=bgo
ARG REPO_NAME=go-template

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /go/src/github.com/omnisinc/${REPO_NAME}
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal

RUN go build -v -tags netgo -ldflags '-extldflags "-static"' -i -o ${SERVICE_NAME} ./cmd/main.go
RUN mv ${SERVICE_NAME} /go/bin/
RUN chmod +x /go/bin/bgo

ENTRYPOINT [ "bgo" ]
