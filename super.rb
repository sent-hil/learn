require 'pry'

class A
  def hello(&blk)
    'A'
  end

  def hi(&blk)
    blk.call
  end
end

class B < A
  def hello
    super
  end

  def hi
    super do
      'B'
    end
  end
end

describe B do
  it 'returns parent result' do
    described_class.new.hello.should == 'A'
  end

  it 'returns child block' do
    described_class.new.hi.should == 'B'
  end
end
