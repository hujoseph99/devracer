const mongoose = require("mongoose");
const Schema = mongoose.Schema;
const schemaTypes = mongoose.Schema.Types;

const UserSchema = new Schema({
  username: {
    type: schemaTypes.String,
    required: true
  },
  nickname: {
    type: schemaTypes.String,
    required: true
  },
  wpm: {
    type: schemaTypes.Decimal128,
    default: 0
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

module.exports = User = mongoose.model("users", UserSchema);
