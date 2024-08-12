# Grafana Mimir
[Documentation](https://grafana.com/docs/mimir)

Grafana Labs metrics backend

## Environments Variables
With example values
```
- MIMIR_MINIO_ACCESS_KEY_ID=zcmapFr9j9nFdgNaCMIU
- MIMIR_MINIO_SECRET_ACCESS_KEY=bLgg5S6BQIeHUvW36ufrCdfKoDnyVf1i6cXIzSqv
- MIMIR_MINIO_BLOCKS_BUCKET_NAME=mimirblocks
- MIMIR_MINIO_ALERT_BUCKET_NAME=mimiralert
- MIMIR_MINIO_RULER_BUCKET_NAME=mimirruler
```

## PromQL
```bash
http_requests_total{job="prometheus",group="canary"}
```
