var _ = require("lodash");
var Dog = require("../models/dog.js");

module.exports = function (app) {
  app.post("/dog", function (req, res) {
    var newDog = new Dog(req.body);
    newDog.save(function (err) {
      if (err) {
        res.json({ info: "error during do create", error: err });
      }
      res.json({ info: "dog created successfully" });
    });
  });

  app.get("/dog", function (req, res) {
    Dog.find(function (err, dogs) {
      if (err) {
        res.json({ info: "error during find dogs", error: err });
      }
      res.json({ info: "dogs found successfully" });
    });
  });

  app.get("/dogs/:id", function (req, res) {
    Dog.findById(req.params.id, function (err, dog) {
      if (err) {
        res.json({ info: "error during find dog", error: err });
      }

      if (dog) {
        res.json({ info: "dog found successfully" });
      } else {
        res.json({ info: "dog not found" });
      }
    });
  });

  app.put("/dog/:id", function (req, res) {
    Dog.findById(req.params.id, function (err, dog) {
      if (err) {
        res.json({ info: "err during find dog", error: err });
      }

      if (dog) {
        _.merge(dog, req.body);
        dog.save(function (err) {
          if (err) {
            res.json({ info: "error during dog create", error: err });
          }
          res.json({ info: "dog updated successfully" });
        });
      } else {
        res.json({ info: "dog not found" });
      }
    });
  });

  app.delete("/dog/:id", function (req, res) {
    Dog.findByIdAndRemove(req.params.id, function (err) {
      if (err) {
        res.json({ info: "error during remove dog", error: err });
      }
      res.json({ info: "dog removed successfully" });
    });
  });
};
