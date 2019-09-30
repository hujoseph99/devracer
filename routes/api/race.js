const express = require("express");
const uuid = require("uuid");
const config = require("config");
const router = express.Router();

const races = [
  {
    id: uuid(),
    snippet:
      "There are many other ways of programming the switching sequence, which will not be discussed here. In many cases the switching sequence is under computer control. In the simple method used above, the intention is to show the fundamental principle upon which pulse width modulation is used."
  },
  {
    id: uuid(),
    snippet:
      "They say there are no stupid questions. That's obviously wrong; I think my question about hard and soft things, for example, is pretty stupid. But it turns out that trying to thoroughly answer a stupid question can take you to some pretty interesting places."
  },
  {
    id: uuid(),
    snippet:
      "Rows and flows of angel hair, and ice cream castles in the air, and feather canyons everywhere. I've looked at clouds that way, but now they only block the sun. They rain and snow on everyone. So many things I would have done, but clouds got in my way."
  },
  {
    id: uuid(),
    snippet:
      "I know you feel where I'm coming from, regardless of the things in my past that I've done. Most of it really was for the hell of the fun, on the carousel so around I spun. With no directions just tryna get some. Tryna chase skirts, living in the summer sun. And so I lost more than I had ever won, and honestly, I ended up with none."
  },
  {
    id: uuid(),
    snippet:
      "All I'm asking is that you do the minimal amount of work in this class to give yourself the illusion that you're actually learning something, and to give me a modicum of self respect like I'm actually teaching a class."
  },
  {
    id: uuid(),
    snippet:
      'The Astonishing Hypothesis it that "You", your joys and your sorrows, your memories and your ambitions, your sense of personal identity and free will, are in fact no more than the behaviour of a vast assembly of nerve cells and their associated molecules. As Lewis Carroll\'s Alice might have phrased it: "You\'re nothing but a pack of neurons."'
  }
];

router.get("/", (req, res) => {
  res.json(races[Math.floor(Math.random() * 6)]);
});

module.exports = router;
