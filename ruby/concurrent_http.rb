require 'concurrent'

require "net/http"
require "uri"

puts "Start #{Time.now.to_i}"

def get
  uri      = URI.parse("http://google.com/")
  http     = Net::HTTP.new(uri.host, uri.port)
  request  = Net::HTTP::Get.new(uri.request_uri)
  response = http.request(request)

  puts "      #{Time.now.to_i}"
end

10.times do
  Concurrent::Promise.new{get()}.then{get()}.execute
end

puts "Queue #{Time.now.to_i}"

sleep(2)

puts "End   #{Time.now.to_i}"
