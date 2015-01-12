require 'elasticsearch-ruby'
require 'pry'

query = {
  "filter" => {
    "query" => {
      "query_string" => {
        "query" => "skyb"
      }
    }
  }
}

client = Elasticsearch::Client.new log: true
resp = client.search index: 'pagehits*', body: query

puts resp['hits']['hits']

# resp.pry
