# cExporter - container stats prometheus exporter

> Like cAdvisor, but more lightweight

![cheerlesscloud/cexporter at docker hub](http://dockeri.co/image/cheerlesscloud/cexporter)

[![Build Status](https://travis-ci.org/CheerlessCloud/cexporter.svg?branch=master)](https://travis-ci.org/CheerlessCloud/cexporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/cheerlesscloud/cexporter)](https://goreportcard.com/report/github.com/cheerlesscloud/cexporter)

### Start in docker

```bash
docker run \
  -d \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --log-opt max-size=10m \
  --name cexporter \
  -p 9167:9167 \
  cheerlesscloud/cexporter:1.0.1
```

### Grafana dashboard

Install from [grafana.com](https://grafana.com/dashboards/9168) or manually [from grafana-dashboard.json](./grafana-dashboard.json).

![grafana dashboard example screenshot](grafana-dashboard-example.png)
