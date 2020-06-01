FROM golang:latest

RUN git clone https://github.com/jmunozal/goshowroom.git /temporal

WORKDIR /temporal/goshowroom/helloworld-server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /temporal/goshowroom/helloworld-server