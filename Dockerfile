FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Gopad CLI" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

ENTRYPOINT ["/usr/bin/gopad-cli"]
CMD ["help"]

RUN apk add --no-cache ca-certificates mailcap bash

COPY dist/binaries/gopad-cli-*-linux-amd64 /usr/bin/
