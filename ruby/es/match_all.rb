require 'elasticsearch-ruby'
require 'pry'

query = {
  "query"=> {
    "match_all" => { }
  }
}

client = Elasticsearch::Client.new log: true
resp = client.search index: 'pagehits*', body: query

puts resp['hits']['hits']

resp.pry
