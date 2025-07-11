// const { unlink } = require("node:fs/promises");

// (async function (path) {
//   try {
//     await unlink(path);
//     console.log(`successfully deleted ${path}`);
//   } catch (error) {
//     console.error("there was on error: ", error.message);
//   }
// })("./cluster.js");

// const { unlink } = require("node:fs");

// unlink("/temp/hello", (err) => {
//   if (err) throw err;
//   console.log("successfully delete /tmp/hello");
// });

// const { unlinkSync } = require("node:fs");

// try {
//   unlinkSync("/tmp/hello");
//   console.log("successfully deleted /tmp/hello");
// } catch (err) {
//   console.log("there was an error: ", err.message);
// }

// import { open } from "node:fs/promises";
const { open } = require("node:fs/promises");
let filehandle;

(async function () {
  try {
    filehandle = await open("thefile.text", "r");
  } catch (err) {
    console.log("there was no file: ", err.message);
  } finally {
    await filehandle?.close();
  }
});
