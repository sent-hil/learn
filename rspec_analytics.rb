require 'rspec'

class Post
  def hi
    'hi'
  end
end

describe Post do
  xit do
    described_class.new.hi.should == 'he'
  end

  it do
    Post.new.should_not be_a(Post)
  end
end
