# Can I've initialize in a module?
#
# Yes

object_space = Object.constants

RSpec.configure do |config|
  config.after(:all) do
    # remove all user defined constants from object space
    (Object.constants-object_space).each do |const|
      Object.send(:remove_const, const)
    end
  end
end

module A
  def initialize
    @yes = 'yes'
  end
end

class Post
  attr_accessor :yes

  include A
end

describe Post do
  it 'have an initialize in module' do
    described_class.new.yes.should == 'yes'
  end
end

# How to convert hash params to getter/setter vars?

class Blog
  def initialize(hash={})
    hash.each do |key, value|
      self.class.instance_eval do
        define_method(key.to_sym) do
          instance_variable_get("@#{key}")
        end

        define_method("#{key.to_sym}=") do |param|
          instance_variable_set("@#{key}", param)
        end
      end

      send("#{key}=", value)
    end
  end
end

describe Blog, 'via define_method' do
  subject do
    described_class.new(:one => 1)
  end

  it 'defines getters from #initialize params' do
    subject.one.should == 1
  end

  it 'defines setters from #initialize params' do
    subject.one = 2
    subject.one.should == 2
  end
end

# How to do above using attr_accessor?

class Blog
  def initialize(hash={})
    self.class.send(:attr_accessor, *hash.keys)

    hash.each do |key, value|
      #self.class.send(:attr_accessor, key.to_sym)

      send("#{key}=", value)
    end
  end
end

describe Blog, 'via attr_accessor' do
  subject do
    described_class.new(:one => 1)
  end

  it 'defines getters from #initialize params' do
    subject.one.should == 1
  end

  it 'defines setters from #initialize params' do
    subject.one = 2
    subject.one.should == 2
  end
end

# Is def get(*params, config={}) valid?
#
# No

#class Post
  #def get(*params, config={})
  #end
#end

#Post.new.get

# What happens when you extend module with #initialize

module B
  def initialize
    @one = 'set in initialize'
  end

  def get
  end
end

class Blogg
  attr_accessor :one

  extend B
end

describe Blogg do
  subject do
    described_class.new
  end

  it do
    subject.one.should == 'set in initialize'
  end
end
