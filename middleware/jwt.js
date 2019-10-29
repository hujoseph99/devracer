const config = require("config");
const jwt = require("jsonwebtoken");

// all requests that need validation must pass along jsonwebtoken in x-auth-token header
const jwtMiddleware = (req, res, next) => {
  const token = req.header("x-auth-token");

  if (!token)
    return res.status(401).json({ msg: "No token, authorization denied" });

  jwt.verify(token, config.get("jwtSecret"), (err, decoded) => {
    if (err)
      return res
        .status(400)
        .json({ msg: "Invalid token, authorization denied" });

    req.userID = decoded;
    next();
  });
};

module.exports = jwtMiddleware;
