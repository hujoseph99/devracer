const express = require("express");
const router = express.Router();
const getRandomSnippet = require("../../multiplayer/snippets");

router.get("/", (req, res) => {
  // var room = require("../../multiplayer/multiplayer").getRoom();
  getRandomSnippet((err, snippet) => {
    if (err) console.log(err);

    // room.room.snippet = snippet.snippet;

    res.json({
      _id: snippet._id,
      snippet: snippet.snippet,
      roomNum: 42
    });
  })
  return;

  // have to first find an open room to figure out what snippet to return
  // if snippet already exists within that room, just return it
  // otherwise, have to get a random snippet from db return that and add the snippet
  //  to the room
  if (room["room"].hasOwnProperty("snippet")) {
    res.json({ roomNum: room.roomNum, snippet: room.room.snippet });
  } else {
    getRandomSnippet
      .then(snippet => {
        room.room.snippet = snippet.snippet;
        res.json({
          _id: snippet._id,
          snippet: snippet.snippet,
          roomNum: room.roomNum
        });
      })
  }
});

module.exports = router;
