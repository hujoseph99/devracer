const express = require("express");
const mongoose = require("mongoose");
const config = require("config");

const app = express();

app.use(express.json());

mongoose
  .connect(config.get("dbUri"), { useNewUrlParser: true })
  .then(() => console.log("db connected..."))
  .catch(err => console.log(err));

const PORT = process.env.PORT || 5000;

app.listen(PORT);
