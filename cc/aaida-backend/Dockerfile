FROM golang:alpine

RUN apk add make

ADD . /go/src/
WORKDIR /go/src/
RUN make

FROM alpine

COPY --from=0 /go/src/dist/aaida /go/bin/aaida

CMD /go/bin/aaida
