FROM golang:1.9

WORKDIR /go/src/github.com/resin-io-projects/resin-go-hello-world

COPY main.go main.go

RUN go get github.com/streadway/amqp/...
RUN go build -o main .
RUN ls 

CMD ./main