FROM golang:1.19-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
RUN go build -o hnService

FROM golang:1.19-alpine
COPY --from=builder /build/hnService /app/
WORKDIR /app
CMD ["./hnService"]