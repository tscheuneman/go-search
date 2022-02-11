FROM golang:1.17-alpine
ENV enviorment = ${ENV}

RUN apk update && apk add git

COPY ./Webserver/server /app
WORKDIR /app

RUN go mod download


RUN go build -o search .


EXPOSE 80
CMD ["/app/search"]