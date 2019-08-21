FROM golang:1.11-alpine3.7 AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh musl-dev gcc

ADD . /src

RUN cd /src && \
 go build -o cexporter .

FROM alpine:3.7

RUN apk add curl

ARG version

LABEL maintainer="CheeressCloud <nnsceg@gmail.com>"
LABEL version=${version}

COPY --from=builder /src/cexporter /usr/bin/cexporter

ENV LOG_FORMAT=json
ENV LOG_LEVEL=info
ENV HTTP_ADDR="0.0.0.0:9167"

EXPOSE 9167

HEALTHCHECK --interval=10s --timeout=3s --start-period=4s \
  CMD curl -f http://${HTTP_ADDR}/metrics || exit 1

CMD [ "/usr/bin/cexporter" ]
