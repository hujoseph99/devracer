const express = require("express");
const router = express.Router();
const bcrypt = require("bcryptjs");
const config = require("config");
const jwt = require("jsonwebtoken");
// const auth = require("../../middleware/auth");

// User Model
const User = require("../../models/User");

// @route   POST api/auth/login
// @desc    Auth user
// @access  Public
// expects 	{ username, password }
router.post("/login", (req, res) => {
  const { username, password } = req.body;

  // Very simple validation of username and password
  // TODO: Improve validation
  if (!username || !password) {
    return res.status(400).json({ msg: "Please enter all fields" });
  }

  // Check for existing user
  User.findOne({ username }).then(user => {
    if (!user) return res.status(400).json({ msg: "User Does not exist" });

    // Validate password
    bcrypt.compare(password, user.password).then(isMatch => {
      if (!isMatch) return res.status(400).json({ msg: "Invalid credentials" });

      // return user details along with a jsonwebtoken
      jwt.sign(
        { id: user._id },
        config.get("jwtSecret"),
        { expiresIn: "1d" },
        (err, token) => {
          if (err) throw err;
          res.json({
            token,
            user: {
              _id: user._id,
              username: user.username,
              wpm: user.wpm
            }
          });
        }
      );
    });
  });
});

// @route   POST api/auth/signup
// @desc    Register new user
// @access  Public
// expects 	{username, email, password}
router.post("/signup", (req, res) => {
  const { username, email, password } = req.body;

  // very simple validation for username, email, and password
  // TODO: Improve validation
  if (!username || !email || !password) {
    return res.status(400).json({ msg: "Please enter all fields" });
  }

  // check if username has already been taken
  User.findOne({ username }).then(user => {
    if (user)
      return res.status(400).json({ msg: "That username is unavailable" });

    // create new user.  register_date and wpm populated by default
    const newUser = new User({
      username,
      email,
      password
    });

    // encrypt password
    bcrypt.genSalt(10, (err, salt) => {
      if (err) throw err;

      bcrypt.hash(newUser.password, salt, (err, hash) => {
        if (err) throw err;

        // save new user to db
        newUser.password = hash;
        newUser.save().then(user => {
          // return user details along with jsonwebtoken
          jwt.sign(
            { id: user._id },
            config.get("jwtSecret"),
            { expiresIn: "1d" },
            (err, token) => {
              if (err) throw err;
              res.json({
                token,
                user: {
                  _id: user._id,
                  username: user.username,
                  wpm: user.wpm
                }
              });
            }
          );
        });
      });
    });
  });
});

module.exports = router;
