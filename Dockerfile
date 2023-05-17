FROM golang:1.19-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -v ./cmd/blitz


FROM debian:buster-slim
COPY --from=builder /app/blitz /app/blitz
CMD ["/app/blitz"]
