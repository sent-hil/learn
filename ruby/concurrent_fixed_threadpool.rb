require 'concurrent'
require 'httparty'
require 'pry'

puts "Start: #{Time.now}"

count = 0

class Ticker
  def update(time, value, reason)
    count += 1
  end
end
ticker = Ticker.new

future = Concurrent::Future.new do
  puts "Koding: #{Time.now}"
  HTTParty.get("https://koding.com")
  puts "Koding: #{Time.now}"
end

future.add_observer(ticker)
future.execute

future = Concurrent::Future.new do
  puts "Google: #{Time.now}"
  HTTParty.get("https://google.com")
  puts "Google: #{Time.now}"
end
future.add_observer(ticker)
future.execute

sleep(5)
puts "Count: #{count}"

puts "End: #{Time.now}"
