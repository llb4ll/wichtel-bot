FROM golang:1.9

WORKDIR /go/src/wichtel
COPY . .

RUN go-wrapper download
RUN go-wrapper install
RUN go test -v .

CMD ["go-wrapper", "run", "-v"]