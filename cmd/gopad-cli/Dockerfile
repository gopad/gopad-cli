FROM --platform=$BUILDPLATFORM golang:1.25.1-alpine3.21 AS builder

RUN apk add --no-cache -U git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/gopad
COPY . /go/src/gopad/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.22

RUN apk add --no-cache ca-certificates mailcap && \
    addgroup -g 1337 gopad && \
    adduser -D -u 1337 -h /var/lib/gopad -G gopad gopad

ENTRYPOINT ["/usr/bin/gopad-cli"]
CMD ["help"]

COPY --from=builder /go/src/gopad/bin/gopad-cli /usr/bin/gopad-cli
WORKDIR /var/lib/gopad
USER gopad
