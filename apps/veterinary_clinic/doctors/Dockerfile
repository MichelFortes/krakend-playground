FROM golang:1.22.5 AS builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o service ./cmd/doctors/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/service /app/service
CMD [ "/app/service" ]
