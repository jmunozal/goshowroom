FROM golang:latest AS build

RUN git clone https://github.com/jmunozal/goshowroom.git /temporal

WORKDIR /temporal/helloworld-server

RUN go build -o /temporal/helloworld-server/server.go

# STAGE 2: Deployment
FROM alpine

USER nobody:nobody
COPY --from=build /temporal/helloworld-server/server.go /server/server

CMD [ "/server/server" ]