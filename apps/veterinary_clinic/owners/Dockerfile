FROM golang:1.22.5 as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o service ./cmd/owners/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/service /app/service
CMD [ "/app/service" ]
