FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN  go mod download

COPY . .
RUN go build -o app

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["/app/app"]
