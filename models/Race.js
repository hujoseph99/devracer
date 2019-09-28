const mongoose = require("mongoose");
const Schema = mongoose.Schema;
const schemaTypes = mongoose.Schema.Types;

const raceSchema = new Schema({
  participants: {
    type: [particpantSchema],
    required: true
  },
  passage: {
    type: schemaTypes.String,
    required: true
  }
});

module.exports = raceSchema;
