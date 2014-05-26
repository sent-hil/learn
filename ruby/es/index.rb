require 'elasticsearch'
require 'jbuilder'
require 'pry'
require 'awesome_print'

client = Elasticsearch::Client.new
client.index  index: 'myindex', type: 'mytype', id: 1, body: { title: 'Test' }

query = Jbuilder.encode do |json|
  json.query do
    json.match do
      json.title do
        json.query 'Test'
      end
    end
  end
end

result = client.search index: 'myindex', type: 'mytype', body: query
pp result['hits']['hits']
