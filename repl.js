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
