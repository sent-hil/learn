require 'elasticsearch'

client = Elasticsearch::Client.new url: "localhost:9200"
entry  = {
  :type      => "redis",
  :body      => {
    :@timestamp => Time.now.strftime("%FT%T%:z")
  }
}

client.index entry
