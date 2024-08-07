# Grafana Alloy
[Documentation](https://grafana.com/docs/alloy)

An Open Telemetry Collector implementation from Grafana Labs

## Environments Variables
With example values
```
- ALLOY_TEMPO_ENDPOINT=http://tempo:4317
- ALLOY_MIMIR_ENDPOINT=http://mimir:9009/api/v1/push
- ALLOY_LOKI_ENDPOINT=http://loki:3100/loki/api/v1/push
```

## Observability

Alloy also sends its own metrics and logs to the appropriate backends
```bash
          [ "Alloy" ] <- self
             /   \
    metrics /     \ logs
           /       \
[ "Mimir" ]         [ "Loki" ]

```
