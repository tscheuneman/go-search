FROM node:12-alpine
ENV enviorment = ${ENV}

WORKDIR /var/www
COPY ./Webserver/package*.json ./

RUN npm install

COPY ./Webserver .

EXPOSE 3010
CMD ["npm", "run", "start"]