FROM golang:1.20-bullseye

COPY src /src
WORKDIR /src

CMD ["go", "run", "main.go"]
