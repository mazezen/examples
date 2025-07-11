const { AsyncResource } = require("node:async_hooks");
const { EventEmitter } = require("node:events");
const path = require("node:path");
const { Worker } = require("node:worker_threads");

const kTaskInfo = Symbol("kTaskInfo");
const kWorkerFreedEvent = Symbol("kWorkerFreedEvent");

class WorkerPoolTaskInfo extends AsyncResource {
  constructor(callback) {
    super("WorkerPoolTaskInfo");
    this.callback = callback;
  }

  done(err, result) {
    this.runInAsyncScope(this.callback, null, err, result);
    this.emitDestroy();
  }
}

class WorkerPool extends EventEmitter {
  constructor(numThreads) {
    super();
    this.numThreads = numThreads;
    this.workers = [];
    this.freedWorkers = [];
    this.tasks = [];

    for (let i = 0; i < numThreads; i++) {
      this.addNewWorker();
    }

    this.on(kWorkerFreedEvent, () => {
      if (this.tasks.length > 0) {
        const { task, callback } = this.tasks.shift();
        this.runTask(task, callback);
      }
    });
  }

  addNewWorker() {
    const worker = new Worker(path.resolve(__dirname, "task_processor.js"));
    worker.on("message", (result) => {
      worker[kTaskInfo].done(null, result);
      worker[kTaskInfo] = null;
      this.freedWorkers.push(worker);
      this.emit(kWorkerFreedEvent);
    });

    worker.on("error", (err) => {
      if (worker[kTaskInfo]) worker[kTaskInfo].done(err, null);
      else this.emit("error", err);

      this.workers.splice(this.workers.indexOf(worker), 1);
      this.addNewWorker();
    });
    this.workers.push(worker);
    this.freedWorkers.push(worker);
    this.emit(kWorkerFreedEvent);
  }

  runTask(task, callback) {
    if (this.freedWorkers.length === 0) {
      this.tasks.push({ task, callback });
      return;
    }

    const worker = this.freedWorkers.pop();
    worker[kTaskInfo] = new WorkerPoolTaskInfo(callback);
    worker.postMessage(task);
  }

  close() {
    for (const worker of this.worker) worker.terminate();
  }
}

module.exports = WorkerPool;
