FROM golang:1.16-alpine3.12 AS builder
RUN apk update && apk add --no-cache git=2.26.3-r0 ca-certificates=20191127-r4

WORKDIR $GOPATH/src/application/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /root/application/app

FROM scratch
LABEL maintainer="iits-consulting.de"
WORKDIR /application/
COPY --from=builder /root/application/ .

ENTRYPOINT ["/application/app"]