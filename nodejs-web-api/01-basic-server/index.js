const http = require("node:http");

http
  .createServer(function (req, res) {
    res.writeHead(200, {
      "content-type": "text/plain",
    });
    res.end("hello world!\n");
  })
  .listen(8003, "localhost");

console.log("Server running at http://127.0.0.1/8003");
