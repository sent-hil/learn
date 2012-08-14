require 'httparty'

def get_events
  HTTParty.get("https://github.com/sent-hil.json")
end
