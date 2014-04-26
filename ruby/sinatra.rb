require 'sinatra'
require 'json'

count = 0

post '*' do
end

post '/fb_redirect' do
   raw = request.body.string
   data = JSON.parse(raw)
   puts "Data: #{data}"

   count += 1
   puts "Count: #{count}"
end

get '/' do
  content_type "application/javascript"
  "hello([]);"
end
