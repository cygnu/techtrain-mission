FROM golang:1.15-alpine as builder
WORKDIR /go/src/bin

RUN apk add --update --no-cache ca-certificates git

ENV GO111MODULE="on"
COPY ../../go.mod .
COPY ../../go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /go/src/bin

FROM scratch as runner

COPY --from=builder /go/src/bin /go/src/app
ENTRYPOINT ["/go/src/app"]