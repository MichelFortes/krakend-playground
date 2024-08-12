# Grafana Loki
[Documentation](https://grafana.com/docs/loki)

Grafana Labs logs backend

## LogQL
```bash
{job="loki.source.gelf.default"} | json | facility="krakend"
```
