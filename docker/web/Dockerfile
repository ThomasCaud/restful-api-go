FROM golang:1.15.2 as builder
RUN apt-get update
RUN apt-get install vim -y
WORKDIR /go/src/api/
COPY . . 
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/main

FROM scratch
COPY --from=builder /go/bin/main /go/bin/main
CMD ["/go/bin/main"]