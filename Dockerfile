FROM alpine:3.7

RUN apk add curl

ARG filename='cexporter_linux_amd64'
ARG version

LABEL maintainer="CheeressCloud <nnsceg@gmail.com>"
LABEL version=${version}

COPY ${filename} /usr/bin/cexporter

ENV LOG_FORMAT=json
ENV LOG_LEVEL=info
ENV HTTP_ADDR="0.0.0.0:9167"

EXPOSE 9167

HEALTHCHECK --interval=10s --timeout=3s --start-period=4s \
  CMD curl -f http://${HTTP_ADDR}/metrics || exit 1

CMD [ "/usr/bin/cexporter" ]