# Apps

## Main Observability Schema

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

## Architecture Components Observability Schema

- [Alloy](alloy/README.md)
