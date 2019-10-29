FROM golang:1.13-alpine AS builder

WORKDIR /workspace/

COPY . .

RUN apk --update --no-cache add ca-certificates && \
  update-ca-certificates && \
  go mod download && \
  go mod verify && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o app .

FROM scratch
COPY --from=gcr.io/berglas/berglas:latest /bin/berglas /bin/berglas
COPY --from=builder /workspace/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV PORT 8080

ENTRYPOINT ["/bin/berglas", "exec", "--", "./app"]
