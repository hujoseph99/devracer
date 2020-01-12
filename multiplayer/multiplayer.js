const getRandomSnippet = require("./snippets");

// expected rooms layout:
// {
//     inProgress: true/false,
//     users: [{
//       username,
//       wpm
//     }]
// }
var room = {
  snippet: "",
  inProgress: false,
  isFinished: true,
  players: []
};

getRandomSnippet((err, snippet) => {
  if (err) console.log(err);

  room.snippet = snippet;
})

// function to deal with websocket logic
startSocket = server => {
  const io = require("socket.io").listen(server);

  io.on("connection", socket => {
    var query = socket.handshake["query"];
    if (query.username) {
      room.players.push({
        username: query.username,
        isParticipant: !room.inProgress,
        wpm: 0
      });
      socket.emit("update", room);
    }
  });
};

// checks rooms and gets an available room number
// returns:
// {
//   room: object corresponding to given roomNum object in "rooms" object above,
//   roomNum: a number
// }
getRoom = () => {
  // for (var prop in rooms) {
  //   if (rooms[prop].hasOwnProperty("isOpen")) {
  //     if (rooms[prop]["isOpen"]) {
  //       return { room: rooms[prop], roomNum: prop };
  //     }
  //   } else {
  //     rooms[prop]["isOpen"] = false;
  //   }
  // }
  // // there is no open room, create a new one -- use random number from 0 - 100
  // let newRoom = Math.floor(Math.random() * 100);
  // rooms[newRoom] = {
  //   isOpen: true
  // };
  // return { room: rooms[newRoom], roomNum: newRoom };
};

module.exports = { startSocket, getRoom };
