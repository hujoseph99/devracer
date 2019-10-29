const express = require("express");
const mongoose = require("mongoose");
const config = require("config");

const app = express();

app.use(express.json());

// Configure routes
app.use("/api/race", require("./routes/api/race"));

mongoose
  .connect(config.get("dbUri"), {
    useNewUrlParser: true,
    useUnifiedTopology: true
  })
  .then(() => console.log("db connected..."))
  .catch(err => console.log(err));

const PORT = process.env.PORT || 5000;

const server = app.listen(PORT);
require("./multiplayer/multiplayer").startSocket(server);
