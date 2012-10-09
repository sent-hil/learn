require 'pry'

#module ApiWrappable
  #def config
    #@config ||= Config.new
  #end

  #def self.define_config_methods(*names)
    #names.each do |name|
      #define_method(name) do |*args|
        #case args.count
        #when 0
          #config.send(name)
        #when 1
          #config.send("#{name}=", args.first)
        #else
          #config.send("#{name}=", args.first.flatten)
        #end
      #end
    #end
  #end

  #define_config_methods :endpoint, :responds_to

  #def get(path)
    #Node.new(config).setup
  #end

  #class Node
    #def initialize(config)
      #@config = config
    #end

    #def setup
      #@url = config.endpoint
    #end
  #end

  #class Config
    #attr_accessor :endpoint, :responds_to
  #end
#end

#describe ApiWrappable do
  #it 'returns nil if not endpoint' do
    #t = Class.new do
      #extend ApiWrappable
    #end

    #t.endpoint.should == nil
  #end

  #it 'sets & returns endpoint' do
    #url = "https://graph.facebook.com"
    #t = Class.new do
      #extend ApiWrappable
      #endpoint url
    #end

    #t.endpoint.should == url
  #end

  #it 'sets & returns :responds_to formats' do
    #formats = [:json, :xml]

    #t = Class.new do
      #extend ApiWrappable
      #responds_to formats
    #end

    #t.responds_to.should == formats
  #end
#end

class Node
  attr_accessor :child, :parent
end

module A
  def get
    node = Node.new
    child = yield if block_given?
    child.parent = node if child

    binding.pry

    node
  end
end

describe A do
  it do
    class T
      extend A

      get do
        puts 1

        get do
          puts 2

          get
        end
      end
    end
  end
end
