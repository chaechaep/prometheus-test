FROM golang:1.19-alpine AS builder

ENV cmd=prometheus-test \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod tidy
RUN go mod download
RUN go build -o ./bin/${cmd} ./api-server/cmd/main.go

FROM alpine:latest AS runtime
ENV cmd=prometheus-test \
    CGO_ENALBED=0 \
    GOOS=linux \
    GOARCH=amd64 \

COPY --from=builder /build/bin/${cmd} /app/bin/

EXPOSE 80
ENTRYPOINT /app/bin/${cmd}