import { createHash, hash } from "node:crypto";
import { readFile } from "node:fs/promises";

const hasher = createHash("sha1");

hasher.setEncoding("hex");
hasher.write(await readFile("package.json"));
hasher.end();

const fileHash = hasher.read();
