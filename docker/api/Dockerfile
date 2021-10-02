FROM golang:1.15.2-alpine as builder
WORKDIR /go/src/github.com/cygnu/techtrain-mission

RUN apk add --update --no-cache ca-certificates git

COPY go.* .
RUN go mod download

COPY . .

ARG CGO_ENABLED=0
RUN go build \
    -o /go/bin/main

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main
ENTRYPOINT ["/app/main"]