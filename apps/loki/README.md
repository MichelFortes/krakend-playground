# Grafana Loki
[Documentation](https://grafana.com/docs/loki)

Solução para Logs da Grafana Labs

## LogQL
```bash
{job="loki.source.gelf.default"} | json | facility="krakend"
```
