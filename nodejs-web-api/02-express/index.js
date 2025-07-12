const express = require("express");

const app = express();

app.get("/", function (req, res) {
  //   res.send("hello world!");
  //   res.send({ name: "json" });
  //   res.status(404).send("sorry, can not find!");
  res.json({ hello: "nodejs" });
});

app.listen(8000, function () {
  console.log("server running at http://127.0.0.1:8000");
});
