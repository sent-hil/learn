require 'elasticsearch-ruby'
require 'pry'

query = {
  "facets" => {
    "username" => {
      "terms" => {
        "field" => "username",
        "size" => 10
      }
    }
  }
}

client = Elasticsearch::Client.new log: true
resp = client.search index: 'pagehits*', body: query

puts resp['facets']['username']['terms']

resp.pry
