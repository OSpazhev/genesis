FROM golang:1.20-bullseye

ADD app/ /go/src/github.com/ospazhev/genesis/app/
WORKDIR /go/src/github.com/ospazhev/genesis/app/

ENV GO111MODULE auto

CMD ["go", "run", "main.go"]
