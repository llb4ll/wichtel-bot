FROM golang:1.21

WORKDIR /go/src/wichtel-bot
COPY src .
COPY config config

RUN go env -w GO111MODULE=auto # only use module when go.mod files is found
RUN go install -v .
RUN go test -v .

CMD ["wichtel-bot"]
