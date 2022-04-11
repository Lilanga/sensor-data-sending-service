#Builder image
FROM golang:alpine3.15 as builder
RUN apk add git
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o sensor-data-sending-service .

#Runner image
FROM alpine:3.15
COPY --from=builder /build/sensor-data-sending-service .

ENTRYPOINT [ "./sensor-data-sending-service" ]