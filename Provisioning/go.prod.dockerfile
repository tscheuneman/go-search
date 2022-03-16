FROM golang:1.17-alpine
ENV enviorment ${ENV}
ENV SERVICE_PORT 5000

RUN apk update && apk add git

COPY ./Webserver/server /app/server

WORKDIR /app/server

RUN go mod download

EXPOSE ${SERVICE_PORT}

RUN go build -o ../search .

ENTRYPOINT /app/search
