require 'concurrent'
require 'open-uri'
require 'pry'

class Ticker
  def get(symbol, year)
    uri = "http://ichart.finance.yahoo.com/table.csv?s=#{symbol}&a=11&b=01&c=#{year}&d=11&e=31&f=#{year}&g=m"
    data = open(uri) {|f| f.collect{|line| line.strip } }
    puts ">>>> #{symbol} #{data[1].split(',')[4].to_f}"
  end
end

price = Concurrent::Future.execute{ Ticker.new.get('TWTR', 2013) }
price.state
price.value
price.state

prices = Concurrent::Promise.new{ Ticker.new.get('AAPL', 2013) }.
           then{ Ticker.new.get('MSFT', 2013) }.
           then{ Ticker.new.get('FORD', 2013) }.
           then{ Ticker.new.get('TSLA', 2013) }.execute
prices.state
sleep(15)

puts "done sleeping"
