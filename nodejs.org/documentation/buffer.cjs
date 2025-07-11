// const { Buffer } = require("node:buffer");

// const buf1 = Buffer.alloc(10);
// console.log(buf1);

// const buf2 = Buffer.alloc(10, 1);
// console.log(buf2);

// const buf3 = Buffer.alloc(10, "hello node", "utf-8");
// console.log(buf3);

// const buf4 = Buffer.allocUnsafe(10);
// console.log(buf4);

// const buf5 = Buffer.from(buf3);
// console.log(buf5);

// const buf = Buffer.from("hello worl1", "utf8");
// console.log(buf.toString("hex"));
// console.log(buf.toString("utf8"));
// console.log(buf.toString("base64"));

// const buf = Buffer.from([1, 2, 3, 4]);
// const uint32Array = new Uint32Array(buf);

// console.log(uint32Array);

// const buf = Buffer.from([1, 2, 3, 4]);
// for (const b of buf) {
//   console.log(b);
// }

// const blod = new Blob(["hello"]);
// blod.bytes().then((butes) => {
//   console.log(butes);
// });

const { Blob } = require("node:buffer");
const { setTimeout: delay } = require("node:timers/promises");

const blob = new Blob(["Hello there"]);

const mc1 = new MessageChannel();
const mc2 = new MessageChannel();

mc1.port1.onmessage = async ({ data }) => {
  console.log(await data.arrayBuffer());
  mc1.port1.close();
};

mc2.port1.onmessage = async ({ data }) => {
  await delay(1000);
  console.log(await data.arrayBuffer());
  mc2.port1.close();
};

mc1.port2.postMessage(blob);
mc2.port2.postMessage(blob);

blob.text().then(console.log);
