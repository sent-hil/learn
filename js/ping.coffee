http = require "http"

watchBroker = (url, restartInterval) ->
  minimumInterval = 400

  console.log("Next interval: ", restartInterval)

  setTimeout ->
    http.get url, (res) ->
      console.log("Got response: ", + res.statusCode)

      watchBroker(url, restartInterval)
    .on 'error', (e) ->
      console.log("Got error: ", e.message)

      restartInterval = if restartInterval <= minimumInterval then minimumInterval else restartInterval/2
      console.log(restartInterval)

      watchBroker(url, restartInterval)
  , 1000

watchBroker("http://asdfgoogle.com", 2000)
