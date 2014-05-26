require 'concurrent'

require "net/http"
require "uri"

puts "Start #{Time.now}"

def get
  uri      = URI.parse("http://google.com/")
  http     = Net::HTTP.new(uri.host, uri.port)
  request  = Net::HTTP::Get.new(uri.request_uri)
  response = http.request(request)

  puts Time.now
end

10.times do
  Concurrent::Promise.new{get()}.then{get()}.execute
end

puts "Queue #{Time.now}"

sleep(2)

puts "End #{Time.now}"
