require "concurrent"

latch = Concurrent::MutexCountDownLatch.new(3)

Thread.new do
  sleep(1)
  latch.count_down
end

latch.wait
