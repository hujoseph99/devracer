const mongoose = require("mongoose");
const Schema = mongoose.Schema;
const schemaTypes = mongoose.Schema.Types;

const RaceSnippetSchema = new Schema({
  snippet: {
    type: schemaTypes.String,
    required: true
  }
});

module.exports = RaceSnippet = mongoose.model(
  "raceSnippets",
  RaceSnippetSchema
);
