import { createReadStream, createWriteStream } from "node:fs";
import { pipeline } from "node:stream/promises";
import { createGzip } from "node:zlib";

await pipeline(
  createReadStream("package.json"),
  createGzip(),
  createWriteStream("package.json")
);
