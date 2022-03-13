FROM golang:1.17-alpine
ENV enviorment = ${ENV}

RUN apk update && apk add git

COPY ./Webserver/server /app/server
COPY ./Webserver/client /app/client

WORKDIR /app/client

# TODO: remove this once once we split off the client into a seperate package

ENV NODE_VERSION=16.13.0
RUN apt install -y curl
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN node --version
RUN npm --version

## END TEMP

RUN npm run build

WORKDIR /app/server

RUN go mod download

EXPOSE 80

RUN go build -o ../search .

ENTRYPOINT /app/search
