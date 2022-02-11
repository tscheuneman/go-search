FROM golang:1.12-alpine
ENV enviorment = ${ENV}

WORKDIR /app/server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./Webserver .

RUN go build -o ./out/search-server .

EXPOSE 8080
CMD ["./out/search-server"]