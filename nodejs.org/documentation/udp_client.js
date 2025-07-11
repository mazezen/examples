// const dgram = require("node:dgram");
// const { Buffer } = require("node:buffer");

// const message = Buffer.from("Some bytes");
// const client = dgram.createSocket("udp4");
// client.send("I'm udp client", 41234, "localhost", (err) => {
//   if (err) {
//     console.log("send message error: ", err);
//   }
//   client.close();
// });

const dgram = require("node:dgram");
const { Buffer } = require("node:buffer");

const buf1 = Buffer.from("Some ");
const buf2 = Buffer.from("bytes");

const client = dgram.createSocket("udp4");
// client.send([buf1, buf2], 41234, "localhost", (err) => {
//   client.close();
// });

client.connect(41234, "localhost", (err) => {
  client.send([buf1, buf2], (err) => {
    client.close();
  });
});
