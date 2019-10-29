const express = require("express");
const router = express.Router();
const RaceSnippet = require("../../models/RaceSnippet");

// TODO: Need to ensure that all members of the same lobby have the same snippet
router.get("/", (req, res) => {
  var room = require("../../multiplayer/multiplayer").getRoom();

  if (room["room"].hasOwnProperty("snippet")) {
    res.json({ roomNum: room.roomNum, snippet: room.room.snippet });
  } else {
    RaceSnippet.countDocuments({}, (err, count) => {
      if (err) console.log(err);

      var rand = Math.floor(Math.random() * count);
      RaceSnippet.findOne()
        .skip(rand)
        .exec((err, snippet) => {
          if (err) console.log(err);

          room.room.snippet = snippet.snippet;
          res.json({
            _id: snippet._id,
            snippet: snippet.snippet,
            roomNum: room.roomNum
          });
        });
    });
  }
});

module.exports = router;
