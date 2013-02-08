var server = require('restify').createServer({
  name: 'foo'
});

server.get('/hello/:name', function echo(req, res, next){
  res.send('hello ' + req.params.name);
  return next();
});

server.listen(8080, function () {
  console.log('Listening at ', server.name, server.url);
});
