require 'elasticsearch-ruby'
require 'pry'

query = {
  "query" => {
    "filtered" => {
      "query" => {
        "query_string" => {
          "query" => "skyb"
        }
      },
      "filter" => {
        "exists" => { "field" => "username" }
      }
    }
  }
}

client = Elasticsearch::Client.new log: true
resp = client.search index: 'pagehits*', body: query

puts resp['hits']['hits']
