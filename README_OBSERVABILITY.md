# Observability

## KrakenD

KrakenD sends
- Traces and Metrics using OTLP
- Logs using GELF

Alloy separate signals and send then to the appropriate backends

```bash
           [ "KrakenD" ]
                 |
                 | traces, metrics and logs
                 |
            [ "Alloy" ]  <- collector
             /   |   \
    metrics /   logs  \ traces
           /     |     \
[ "Mimir" ]   [ "Loki" ]  [ "Tempo" ]
```

## Alloy

Alloy also sends its own metrics and logs to the appropriate backends
```bash
          [ "Alloy" ] <- self
             /   \
    metrics /     \ logs
           /       \
[ "Mimir" ]         [ "Loki" ]

```
