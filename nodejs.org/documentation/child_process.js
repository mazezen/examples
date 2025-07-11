// const { spawn } = require("node:child_process");
// const ls = spawn("ls", ["-lh", "/usr"]);

// ls.stdout.on("data", (data) => {
//   console.log(`stdout: ${data}`);
// });

// ls.stderr.on("data", (data) => {
//   console.error(`stderr: ${data}`);
// });

// ls.on("close", (code) => {
//   console.log(`child process exited with code ${code}`);
// });

// const { exec } = require("node:child_process");
// exec('"./test.sh');
// exec('echo "The \\$HOME variable is $HOME"');

// const { exec } = require("node:child_process");
// exec("cat *.js missing_file | wc -l", (error, stdout, stderr) => {
//   if (error) {
//     console.error(`exec error: ${error}`);
//     return;
//   }

//   console.log(`stdout: ${stdout}`);
//   console.log(`stderr: ${stderr}`);
// });

// const util = require("node:util");
// const exec = util.promisify(require("node:child_process").exec);

// async function lsExample() {
//   const { stdout, stderr } = await exec("ls");
//   console.log("stdout: ", stdout);
//   console.error("stderr:", stderr);
// }
// lsExample();

// const { exec } = require("node:child_process");
// const { error } = require("node:console");
// const controller = new AbortController();
// const { signal } = controller;
// const child = exec("grep ssh", { signal }, (error) => {
//   console.error(error);
// });
// controller.abort();

// const { execFile } = require("node:child_process");
// const { error } = require("node:console");
// const { stderr } = require("node:process");
// const child = execFile("node", ["--version"], (error, stdout, stderr) => {
//   if (error) {
//     throw error;
//   }
//   console.log(stdout);
// });

// const child2 = execFile("node", ["-v"], (error, stdout, stderr) => {
//   if (error) {
//     throw error;
//   }

//   console.log(stdout);
// });

// const util = require("node:util");
// const execFile = util.promisify(require("node:child_process").execFile);
// async function getVersion() {
//   const { stdout } = await execFile("node", ["--version"]);
//   console.log(stdout);
// }
// getVersion();

// const { execFile } = require("node:child_process");
// const controller = new AbortController();
// const { signal } = controller;
// const child = execFile("node", ["--version"], { signal }, (error) => {
//   console.log(error);
// });
// controller.abort();

// const { fork } = require("node:child_process");
// const process = require("node:process");

// if (process.argv[2] === "child") {
//   setTimeout(() => {
//     console.log(`Hello from ${process.argv[2]}!`);
//   }, 1_000);
// } else {
//   const controller = new AbortController();
//   const { signal } = controller;
//   const child = fork(__filename, ["child"], { signal });
//   child.on("error", (err) => {});
//   controller.abort();
// }

// const { spawn } = require("node:child_process");
// const ps = spawn("ps", ["ax"]);
// const grep = spawn("grep", ["ssh"]);

// ps.stdout.on("data", (data) => {
//   grep.stdin.write(data);
// });

// ps.stderr.on("data", (data) => {
//   console.error(`ps stderr: ${data}`);
// });

// ps.on("close", (code) => {
//   if (code != 0) {
//     console.log(`ps process exited with code ${code}`);
//   }
//   grep.stdin.end();
// });

// grep.stdout.on("data", (data) => {
//   console.log(data.toString());
// });

// grep.stderr.on("data", (data) => {
//   console.error(`grep stderr: ${data}`);
// });

// grep.on("close", (code) => {
//   if (code !== 0) {
//     console.log(`grep process exited with code ${code}`);
//   }
// });

// const subprocess = spawn("bad_command");
// subprocess.on("error", (err) => {
//   console.err("Failed to start subprocess.");
// });

// const { spawn } = require("node:child_process");
// const controller = new AbortController();
// const { signal } = controller;
// const grep = spawn("grep", ["ssh"], { signal });
// grep.on("error", (err) => {});
// controller.abort();

const { openSync } = require("node:fs");
const { spawn } = require("node:child_process");
const out = openSync("./out.log", "a");
const err = openSync("./out.log", "a");

const subprocess = spawn("prg", [], {
  detached: true,
  stdio: ["ignore", out, err],
});

subprocess.unref();
