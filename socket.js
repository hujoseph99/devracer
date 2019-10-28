var rooms = {};

startSocket = server => {
  const io = require("socket.io").listen(server);

  io.on("connection", socket => {
    var query = socket.handshake["query"];

    if (rooms.hasOwnProperty(query.room)) {
      rooms[query.room][query.username] = query.wpm;
    } else {
      rooms[query.room] = {};
      rooms[query.room][query.username] = query.wpm;
    }

    console.log(rooms);
  });
};

module.exports = startSocket;
