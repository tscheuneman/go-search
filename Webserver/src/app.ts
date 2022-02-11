import express from "express";

import moment from "moment";

require('dotenv').config();
const bodyParser = require('body-parser');
const app = express();

app.set('trust proxy', true);

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
  extended: true
}));

app.get('/', (req: any, res) => {
  res.send('Welcome to the ASU Print and Imaging Lab API HEHE');
});

if(!module.parent) {
  app.listen(3010, () => {
    console.log('Dev server initated');
  });
}


module.exports = app;