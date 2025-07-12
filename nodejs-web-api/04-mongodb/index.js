const express = require("express");
const app = express();

const bodyParser = require("body-parser");
const mongoose = require("mongoose");
mongoose.connect("mongodb://localhost", {
  user: "root",
  pass: "123456",
});

app.use(bodyParser.json());
app.use(
  bodyParser.urlencoded({
    extended: true,
  })
);

require("./cat_routes.js")(app);

app.listen(3000, function () {
  console.log("server running at http://127.0.0.1:3000");
});
