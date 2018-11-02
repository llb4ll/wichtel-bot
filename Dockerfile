FROM golang:1.11

WORKDIR /go/src/wichtel-bot
COPY src .
COPY config config

RUN go get -d -v .
RUN go test -v .
RUN go install -v .

CMD ["wichtel-bot"]
