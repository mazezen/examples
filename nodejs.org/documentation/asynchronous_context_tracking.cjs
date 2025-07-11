const http = require("node:http");
const { AsyncLocalStorage } = require("node:async_hooks");

// const asyncLocalStorage = new AsyncLocalStorage();

// function logWithId(msg) {
//   const id = asyncLocalStorage.getStore();
//   console.log(`${id !== undefined ? id : "-"}:`, msg);
// }

// let idSeq = 0;
// http
//   .createServer((req, res) => {
//     asyncLocalStorage.run(idSeq++, () => {
//       logWithId("start");

//       setImmediate(() => {
//         logWithId("finish");
//         res.end();
//       });
//     });
//   })
//   .listen(8081);

// http.get("http://localhost:8081");
// http.get("http://localhost:8081");

// const asyncLocalStorage = new AsyncLocalStorage();
// const runInAsyncScope = asyncLocalStorage.run(123, () =>
//   AsyncLocalStorage.snapshot()
// );
// const result = asyncLocalStorage.run(321, () =>
//   runInAsyncScope(() => asyncLocalStorage.getStore())
// );
// console.log(result);

const asyncLocalStorage = new AsyncLocalStorage();

class Foo {
  #runInAsyncScope = AsyncLocalStorage.snapshot();

  get() {
    return this.#runInAsyncScope(() => asyncLocalStorage.getStore());
  }
}

const foo = asyncLocalStorage.run(123, () => new Foo());
console.log(asyncLocalStorage.run(321, () => foo.get()));
