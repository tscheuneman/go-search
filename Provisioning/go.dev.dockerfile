FROM golang:1.17-alpine
ENV enviorment = ${ENV}

RUN apk update && apk add git

COPY ./Webserver/server /app/server
COPY ./Webserver/client/build /app/client

WORKDIR /app/server

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 80

ENTRYPOINT CompileDaemon --build="go build -o ../search ." --command="/app/search"
