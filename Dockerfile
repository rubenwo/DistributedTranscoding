FROM golang:1.16 AS builder

WORKDIR /go/go-encoder

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-encoder ./cmd/encoder

FROM alpine:3.12
RUN apk --no-cache add ca-certificates ffmpeg

WORKDIR /root/
COPY --from=builder /go/go-encoder/go-encoder .
COPY --from=builder /go/go-encoder/assets/ ./assets/

CMD [ "./go-encoder" ]