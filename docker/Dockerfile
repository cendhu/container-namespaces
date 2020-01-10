FROM ubuntu:latest

RUN apt-get update

RUN mkdir /persistData

WORKDIR /app

ENV TEST_ENV_VAR="This is a winter school demo"

VOLUME /persistData

COPY test.sh /app/test.sh

CMD ["/app/test.sh"]
