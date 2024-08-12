# Grafana
| [Documentation](https://grafana.com/oss/grafana/)  |  [Interface](http://localhost:3000) |

An observability tool for visualizing telemetry data.

## Dashboards

Comig soon

### Extraindo Métricas de Traces no Grafana (TraceQL Metrics)

https://github.com/grafana/tempo/blob/main/docs/design-proposals/2023-11%20TraceQL%20Metrics.md

Exemplo de query:
```
{ status = ok } | rate() by (span.http.route)
```
