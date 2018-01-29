# cExporter - container stats Prometheus exporter.
## Like cAdvisor, but more lightweight

![cheerlesscloud/cexporter at docker hub](http://dockeri.co/image/cheerlesscloud/cexporter)

[![Build Status](https://travis-ci.org/CheerlessCloud/cexporter.svg?branch=master)](https://travis-ci.org/CheerlessCloud/cexporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/cheerlesscloud/cexporter)](https://goreportcard.com/report/github.com/cheerlesscloud/cexporter)

### Start in docker
```bash
docker run \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -p 9167:9167 \
  --log-opt max-size=10m \
  -d --name cexporter \
  cheerlesscloud/cexporter
```
