# Grafana Loki
[Documentation](https://grafana.com/docs/loki)

Grafana Labs logs backend

## Environments Variables
With example values
```
- LOKI_MINIO_ACCESS_KEY_ID=ddLKSQA8wkQLKX7uYdIK
- LOKI_MINIO_SECRET_ACCESS_KEY=5T5o1Gq7ha3xyfrUlPbYfrVAWywfZKlKePI4sUzq
- LOKI_MINIO_BUCKET_NAME=loki
- LOKI_MINIO_REGION=local
```

## LogQL
```bash
{job="loki.source.gelf.default"} | json | facility="krakend"
```
