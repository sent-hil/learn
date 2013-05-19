require 'sinatra'

delete '/cocacola' do
  puts 'coca cola'
end

get '/' do
  content_type "application/javascript"
  "hello([]);"
end
