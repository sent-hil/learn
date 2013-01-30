module RiverApp
  module Middlewares
    class RouteParams
      def initialize(app, url)
        @app = app
        @url = url
      end

      def call(env)
        split_url_params = env['REQUEST_URI'].split('/')
        split_url_params.delete('')

        url = @url.split('/')
        url.delete('')

        url.each_with_index do |key, index|
          next unless key.include?(':')
          env['params'][key.gsub(':','')] = split_url_params[index]
        end

        @app.call(env)
      end
    end
  end
end
