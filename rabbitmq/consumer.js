var amqp = require('amqp');

var connection = amqp.createConnection();
connection.on('ready', function() {
  connection.queue('simple', function(q){
    console.log(q.name + ' queue open');
    q.bind("amq.direct", "routing_key");

    q.subscribe(function (message, headers, deliveryInfo) {
      console.log('Message:', message.data.toString('utf-8'));
    });
  });
});
