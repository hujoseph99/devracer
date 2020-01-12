const RaceSnippet = require("../models/RaceSnippet");

const getRandomSnippet = callback => {
  RaceSnippet.countDocuments({}, (err, count) => {

    if (err) return callback(err, null);

    let rand = Math.floor(Math.random() * count);
    RaceSnippet.findOne()
      .skip(rand)
      .exec((err, snippet) => {
        if (err) return callback(err, null);

        return callback(null, snippet);
      });
  });
}

module.exports = getRandomSnippet;
