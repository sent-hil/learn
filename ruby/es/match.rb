require 'elasticsearch-ruby'
require 'pry'

query = {
  "query"=> {
    "match" => { "username" => "skyb" }
  },
  "sort"=> {
    "@timestamp"=> { "order"=> "desc" }
  },
  "size"=> 100
}

client = Elasticsearch::Client.new log: true
resp = client.search index: 'pagehits*', body: query

puts resp['hits']['hits']

# resp.pry
