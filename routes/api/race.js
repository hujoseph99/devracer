const express = require("express");
const router = express.Router();
const RaceSnippet = require("../../models/RaceSnippet");

router.get("/", (req, res) => {
  var room = require("../../multiplayer/multiplayer").getRoom();

  // have to first find an open room to figure out what snippet to return
  // if snippet already exists within that room, just return it
  // otherwise, have to get a random snippet from db return that and add the snippet
  //  to the room
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
