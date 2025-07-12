const bodyParser = require("body-parser");
const express = require("express");
const app = express();

app.use(bodyParser.json());
app.use(
  bodyParser.urlencoded({
    extended: true,
  })
);
const cats = require("./cats.js")(app);

app.listen(8000, function () {
  console.log("server running at http://127.0.0.1:8000");
});
