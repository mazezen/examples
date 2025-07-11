const { createServer } = require("node:http");
const { AsyncResource, executionAsyncId } = require("node:async_hooks");

const server = createServer((req, res) => {
  req.on(
    "close",
    AsyncResource.bind(() => {})
  );

  req.on("close", () => {});
  res.end();
}).listen(3000);
