FROM golang:1.11.3-alpine3.8 as builder

WORKDIR /go/src/github.com/ic2hrmk/azck
COPY . .
RUN apk add git dep && \
    dep ensure -v --vendor-only && \
    CGO_ENABLED=0 go build -a -o azck entry/entry.go

FROM alpine:3.8
WORKDIR /usr/bin/sbin
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/ic2hrmk/azck/azck .
CMD ["./azck"]
