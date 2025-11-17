FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o certsync-cos .

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /build/certsync-cos .

RUN mkdir -p /app/config /app/certs

ENV TZ=Asia/Shanghai

CMD ["/app/certsync-cos", "-config", "/app/config/config.yaml"]