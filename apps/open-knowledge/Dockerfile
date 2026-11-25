# syntax=docker/dockerfile:1

FROM docker.io/library/node:24-trixie-slim@sha256:4d00acb4bcc42c454c4eb801526f7ad7f772ea3bd4fd168c5cc2e384abcbb9d4
ARG VERSION

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
    ca-certificates \
    catatonit \
    git \
    tzdata \
  && rm -rf /var/lib/apt/lists/*

RUN npm install --global "@inkeep/open-knowledge@${VERSION}" \
  && npm cache clean --force

COPY --chmod=755 entrypoint.sh /entrypoint.sh

RUN mkdir /data && chown node:node /data

ENV \
  OK_BIND=0.0.0.0 \
  PORT=8080

USER node:node
WORKDIR /data

ENTRYPOINT ["/usr/bin/catatonit", "--", "/entrypoint.sh"]
