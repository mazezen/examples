"use strict";

// const Database = require("better-sqlite3");
// const db = new Database(":memory:");

// try {
//   db.exec(`
//     CREATE TABLE data(
//         key INTEGER PRIMARY KEY,
//         value TEXT
//     ) STRICT
//     `);

//   const insert = db.prepare(
//     "INSERT INTO data (key, value) VALUES (?, ?)"
//   );
//   insert.run(1, "hello");
//   insert.run(2, "world");

//   const query = db.prepare("SELECT * FROM data ORDER BY key");

//   console.log(query.all());

//   db.close();
// } catch (error) {
//   console.error("数据库操作失败: ", error.message);
// }

const Database = require("better-sqlite3");
const db = new Database("foobar.db", { verbose: console.log });

try {
  db.exec(`DROP TABLE cats`);
  db.exec(`
      CREATE TABLE cats(
          name TEXT,
          age TEXT
      ) STRICT
      `);

  const insert = db.prepare(
    "INSERT INTO cats (name, age) VALUES (@name, @age)"
  );
  const insertMany = db.transaction((cats) => {
    for (const cat of cats) insert.run(cat);
  });

  insertMany([
    { name: "Joery2", age: "2" },
    { name: "Joery4", age: "4" },
    { name: "Joery5", age: "5" },
    { name: "Joery8", age: "8" },
  ]);
} catch (error) {
  console.error("数据库操作失败: ", error.message);
}
