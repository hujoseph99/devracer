var rooms = {};

startSocket = server => {
  const io = require("socket.io").listen(server);

  io.on("connection", socket => {
    var query = socket.handshake["query"];
    if (query.roomNum && query.username && rooms[query.roomNum]) {
      console.log(query.roomNum);
      console.log(rooms);
      // rooms[query.roomNum]["username"] = query.username;
    }
  });
};

getRoom = () => {
  for (var prop in rooms) {
    if (rooms[prop].hasOwnProperty("isOpen")) {
      if (rooms[prop]["isOpen"]) {
        return { room: rooms[prop], roomNum: prop };
      }
    } else {
      rooms[prop]["isOpen"] = false;
    }
  }
  // there is no open room, create a new one -- use random number from 0 - 100
  // TODO: Need to disallow usernames of "isOpen"
  let newRoom = Math.floor(Math.random() * 100);
  rooms[newRoom] = {
    isOpen: true
  };
  return { room: rooms[newRoom], roomNum: newRoom };
};

module.exports = { startSocket, getRoom };
