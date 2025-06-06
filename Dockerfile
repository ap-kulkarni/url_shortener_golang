FROM golang:alpine AS builder
LABEL authors="ameykulkarni"

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./url_shortner

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/url_shortner ./
CMD ["/app/url_shortner"]
