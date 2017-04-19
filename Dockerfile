FROM alpine:3.4
MAINTAINER Sadat Hossain <sadat.hossain@hotmail.com>

COPY ./bin/mitnotes /app/mitnotes
COPY ./public /app/public

WORKDIR /app
CMD ["./mitnotes"]
EXPOSE 3000
