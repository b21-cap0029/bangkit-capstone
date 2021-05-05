FROM golang:alpine

ADD . /go/src/
WORKDIR /go/src/
RUN go build main.go

FROM alpine

COPY --from=0 /go/src/main /go/bin/main

CMD /go/bin/main
