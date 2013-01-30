class Post
  # can't extend self on class
  # extend self

  def hi
    'hi'
  end
end

describe Post do
  it do
    described_class.hi.should == ''
  end
end
