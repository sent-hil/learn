10.times.map do |i|
  Thread.new do
    `git clone https://github.com/sent-hil/es_utils es_utils-#{i}`
  end
end.each(&:join)
