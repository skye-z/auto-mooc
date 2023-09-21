FROM alpine:3.18.3
MAINTAINER skye-z <skai-zhang@hotmail.com>

COPY auto-mooc /usr/local/bin/

RUN addgroup -S nonroot && \
    adduser -S betax -G nonroot && \
    cd /usr/local/bin/ && \
    chmod +x /usr/local/bin/auto-mooc && \
    auto-mooc init

USER betax

EXPOSE 80

ENTRYPOINT [ "auto-mooc" ]