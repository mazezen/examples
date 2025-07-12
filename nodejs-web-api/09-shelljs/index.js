#!/usr/bin/env node

var fs = require("fs");
var sh = require("shelljs");

var file = process.argv.pop();
console.log(file);

fs.readFile(file, { encoding: "utf8" }, function (err, data) {
  console.log("readFile: %s", data);
});

console.log("cat: %s", sh.cat(file));
