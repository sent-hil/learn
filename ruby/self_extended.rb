module A
  class << self
    attr_reader :extendeds
  end

  # Keep track of classes that extendeds this module.
  def self.extended(klass)
    @extendeds ||= []
    @extendeds << klass
  end
end

class B
  extend A
end

puts A.extendeds
