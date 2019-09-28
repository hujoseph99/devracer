const mongoose = require("mongoose");
const Schema = mongoose.Schema;
const schemaTypes = mongoose.Schema.Types;

const userSchema = new Schema({
  name: {
    type: schemaTypes.String,
    required: true
  },
  wpm: {
    type: schemaTypes.Decimal128,
    required: true
  },
  email: {
    type: schemaTypes.String,
    required: true
  },
  password: {
    type: schemaTypes.String,
    required: true
  },
  register_date: {
    type: schemaTypes.Date,
    default: Date.now()
  }
});

module.exports = userSchema;
