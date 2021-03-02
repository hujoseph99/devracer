FROM golang:1.16-alpine

WORKDIR /go/src/hujoseph99/typing

COPY . ./

RUN go get
RUN go build

CMD ["backend"]


