require 'elasticsearch-ruby'
require 'pry'

client = Elasticsearch::Client.new log: true
client.bulk :body => [
  { :index => {
      :_index => 'myindexa', :_type => 'mytype', :_id => '1',
      :data => { :title => 'Test' } }
  },
  { :update => {
      :_index => 'myindexb', :_type => 'mytype', :_id => '1',
      :data => { :doc => { :title => 'Update' } } }
  },
  { :delete => {
      :_index => 'myindexqc', :_type => 'mytypeC', :_id => '1' }
  }
]
