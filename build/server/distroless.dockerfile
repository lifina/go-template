FROM golang:1.14 AS builder

ARG SERVICE_NAME=bgo
ARG REPO_NAME=go-template

WORKDIR /go/src/github.com/omnisinc/${REPO_NAME}
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
COPY cmd cmd
COPY internal internal

RUN go build -v -tags netgo -ldflags '-extldflags "-static"' -i -o ${SERVICE_NAME} ./cmd/main.go
RUN mv ${SERVICE_NAME} /bin/

FROM gcr.io/distroless/static
COPY --from=builder /bin/${SERVICE_NAME} /bin/${SERVICE_NAME}

ENTRYPOINT ["/bin/bgo"]
