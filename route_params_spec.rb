require 'integration/spec_helper'
require 'middlewares/route_params'

describe RiverApp::Middlewares::RouteParams do
  include Goliath::TestHelper

  before(:all) do
    @klass = Class.new(Goliath::API) do
      use Goliath::Rack::Params
      use RiverApp::Middlewares::RouteParams, '/rivers/:id'

      def response(env)
        [200, {}, params]
      end
    end
  end

  subject do
    with_api(@klass) do
      get_request(:path => '/rivers/1') do |api|
        @response = api.response
      end
    end

    @response
  end

  it 'populates params hash with url params' do
    subject.should == ["id", "1"].to_s
  end
end
