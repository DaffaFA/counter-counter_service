FROM golang:1.22.4-alpine AS builder
WORKDIR /app
COPY . .
RUN set -eux; \
  go mod tidy; \
  go build -o server

FROM alpine:3.16 AS runner
WORKDIR /app
COPY --from=builder /app/server ./server
CMD ["/app/server"]
