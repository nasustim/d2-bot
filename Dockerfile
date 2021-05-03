FROM golang:1.16

WORKDIR /go/src/main
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build ./src/main.go

CMD ["./main"]