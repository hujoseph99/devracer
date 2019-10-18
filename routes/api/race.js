const express = require("express");
const router = express.Router();

const RaceSnippet = require("../../models/RaceSnippet");

router.get("/", (req, res) => {
  // console.log(RaceSnippet.collection.collectionName);
  RaceSnippet.countDocuments({}, (err, count) => {
    if (err) console.log(err);

    var rand = Math.floor(Math.random() * count);
    RaceSnippet.findOne()
      .skip(rand)
      .exec((err, snippet) => {
        if (err) console.log(err);
        res.json(snippet);
      });
  });
});

module.exports = router;
