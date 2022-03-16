FROM golang:1.17-alpine
ENV enviorment ${ENV}
ENV SERVICE_PORT 5000

RUN apk update && apk add git

COPY ./Webserver/server /app/server

WORKDIR /app/server

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE ${SERVICE_PORT}

ENTRYPOINT CompileDaemon --build="go build -o ../search ." --command="/app/search"
