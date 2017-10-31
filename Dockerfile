FROM golang:1.9

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install
RUN go test .

CMD ["go-wrapper", "run"]