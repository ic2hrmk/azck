FROM golang:1.11.3-alpine3.8

WORKDIR /go/src/github.com/ic2hrmk/azck
COPY . .

RUN go build -o app entry/entry.go && mv app /go/bin/

CMD ["app"]
