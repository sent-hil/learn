var amqp = require('amqp');

var connection = amqp.createConnection({
  defaultExchangeName: "amq.direct"
});

connection.on('ready', function() {
  console.log("Publishing...");
  connection.publish("simple", "Hello from Koding!");
});
