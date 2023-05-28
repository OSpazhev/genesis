FROM golang:1.20-bullseye

ADD app/ /go/src/github.com/ospazhev/genesis/app/
WORKDIR /go/src/github.com/ospazhev/genesis/app/

ENV GO111MODULE auto

ENV DATAPATH "/data/"
ENV CoinAPIKey "01F76A3F-1B3A-4B61-84A7-2EF25DD8A3CA"

CMD ["go", "run", "main.go"]
