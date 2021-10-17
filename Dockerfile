FROM golang:1.16-alpine

ADD . /jaeminbot
WORKDIR "/jaeminbot"
RUN apk add --no-cache --update bash make git build-base
RUN ["make", "build"]
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["/jaeminbot/jaeminbot -token $TOKEN -store $STORE"]