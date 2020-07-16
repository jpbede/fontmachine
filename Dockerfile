FROM golang:1.14.5-alpine3.12

ADD . /code
WORKDIR /code

RUN go build commands/fontmachined.go
RUN mv fontmachined /

# cleanup
RUN rm -rf /root/.cache && rm -rf /code && rm -rf /go

WORKDIR /

ENV GIN_MODE=release
ENTRYPOINT ["/fontmachined", "-p", "/fonts", "-l", ":8080"]