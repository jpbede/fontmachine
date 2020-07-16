FROM golang:1.14.5-alpine3.12

ADD . /code
WORKDIR /code

RUN go build commands/fontmachined.go
RUN mv fontmachined /

# cleanup
RUN rm -rf /root/.cache && rm -rf /code

WORKDIR /
VOLUME ["/fonts"]
CMD ["/fontmachined", "-p", "/fonts"]