# Does block_given? work with &blk?

require 'pry'

def hello(&blk)
  binding.pry
end

describe do
  it do
    hello do
      puts 'hi'
    end.should == ''
  end

  it do
    hello
  end
end
