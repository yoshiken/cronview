FROM golang:1.11

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/app

CMD ["dep", "ensure"]
