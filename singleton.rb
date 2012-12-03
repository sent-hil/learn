require 'singleton'

class Post
  include Singleton

  def track
    'track'
  end
end

describe Post do
  it do
    # use instance to access singleton
    described_class.instance.track.should == 'track'
  end
end
