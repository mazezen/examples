// const ac = new AbortController();

// ac.signal.addEventListener("abort", () => console.log("Aborted!"), {
//   once: true,
// });
// ac.abort();

// console.log(ac.signal.aborted);

// const ac = new AbortController();

// ac.signal.aborted = () => console.log("aborted!");

// ac.signal.addEventListener(
//   "abort",
//   (event) => {
//     console.log(event.type);
//   },
//   { once: true }
// );

// ac.abort();

// const res = fetch("https://nodejs.org/api/documentation.json");
// if (res.ok) {
//   const data = res.json();
//   console.log(data);
// }
