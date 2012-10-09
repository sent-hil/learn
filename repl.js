var stdout = process.stdout;
var readline = require('readline');

var repl = readline.createInterface({
  input: process.stdin,
  output: stdout
});

repl.setPrompt('> ');
repl.prompt();

process.on("uncaughtException", function (error) {
  stdout.write(error.stack);
});

repl.on('line', function(cmd) {
  console.log(eval(cmd));
});

var net = require('net');
var HOST = '127.0.0.1';
var PORT = 27017;
var connection = new net.Socket();
var BSON = require('bson');
var Buffer = require('buffer').Buffer;

connection.connect(PORT, HOST, function() {
  console.log('connected');

  client = net.createConnection(PORT, HOST)
  console.log(client);

  db   = "sample_db"
  coll = "test"
  docs = [{a:1}]
  msg  = ""
  size = Buffer.byteLength(coll);

  for(var i = 0; i < docs.length; i++) {
    console.log('----')
    console.log(BSON.calculateObjectSize(docs[i]));
    console.log('----')
  };

  console.log(size);

  //result = client.write(msg);
  //console.log(result);
});
