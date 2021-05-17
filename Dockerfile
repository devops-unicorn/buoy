FROM golang:1.16.3 as builder
WORKDIR /go/src/github.com/devops-unicorn/buoy
RUN go get -d -v golang.org/x/net/html
COPY .  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o buoy .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/devops-unicorn/buoy .
# leave this as blackhole until we have real binary file
ENTRYPOINT ["tail", "-f", "/dev/null"]