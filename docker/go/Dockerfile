FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go get github.com/cosmtrek/air@v1.29.0
CMD ["air", "-c", ".air.toml"]