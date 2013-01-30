require 'pry'

class A
  def self.define_method_with_blks(name)
    define_method(name) do |&blk|
      instance_variable_set("@#{name}", blk)
    end
  end

  define_method_with_blks :hello
end

A.new.hello {puts 'hi'} #=> Proc object
