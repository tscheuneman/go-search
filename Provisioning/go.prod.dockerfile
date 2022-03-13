FROM golang:1.17-alpine
ENV enviorment = ${ENV}

RUN apk update && apk add git

COPY ./Webserver/server /app/server
COPY ./Webserver/client /app/client

WORKDIR /app/client

RUN npm run build

WORKDIR /app/server

RUN go mod download

EXPOSE 80

RUN go build -o ../search .

ENTRYPOINT /app/search