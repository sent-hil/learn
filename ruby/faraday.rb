require 'faraday'

conn = Faraday.new(:url => 'https://graph.facebook.com') do |faraday|
  faraday.request  :url_encoded
  faraday.response :logger
  faraday.adapter  Faraday.default_adapter
end

conn.delete 100002930317357_327528810688127 , :access_token => 'AAAFQSimj1rkBAEEzYjDGNyIag6EZBSwYHl2jXDJlmIdAjc3nlXpkajStjqhPyrSLZBUQ8sfV6H40LktIa1zvAILHZCZBgJgPowtiJ2K9lQZDZD'
