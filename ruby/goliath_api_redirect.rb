# Problem: Divide api into version to provide a consistent experience to developers.

require 'goliath'
require 'active_support/all'

module V01
  class Hello
    def self.call(env)
      [200, {}, {'response' => 'V01'}]
    end
  end
end

module V02
  class Hello
    def self.call(env)
      [200, {}, {'response' => 'V02'}]
    end
  end
end

class HelloApp < Goliath::API
  use Goliath::Rack::Params

  def response(env)
    version = env['HTTP_ACCEPTS'].tr('^0-9','')
    version = "V"+version

    version.constantize::Hello.call(env)
  end
end
