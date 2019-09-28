const mongoose = require("mongoose");
const Schema = mongoose.Schema;
const schemaTypes = mongoose.Schema.Types;

const participantSchema = new Schema({
  userID: {
    type: schemaTypes.ObjectId,
    required: true
  },
  wpm: {
    type: schemaTypes.Decimal128,
    required: true
  },
  mistakes: {
    type: [schemaTypes.Number],
    required: true
  }
});

module.exports = participantSchema;
