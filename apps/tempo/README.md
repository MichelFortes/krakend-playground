# Grafana Tempo
[Documentation](https://grafana.com/docs/tempo)

Grafana Labs traces backend

## Environments Variables
With example values
```
- TEMPO_MINIO_ACCESS_KEY_ID=3Gv862RTW4xKjYv6m7Ec
- TEMPO_MINIO_SECRET_ACCESS_KEY=gpye4QVWjyZlBXsBWXyHL7NPmTCOS0JI0BrH1wvW
- TEMPO_MINIO_BUCKET_NAME=tempo
```

## TraceQL
```bash
{} | count() > 2
```
