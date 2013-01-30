require 'open-uri'
require 'octokit'
require 'nokogiri'
require 'twitter'

@client = Octokit::Client.new
@gists  = @client.gists('sent-hil')

def match_hacker_school_entries?(x)
  /\+\d+/.match(x['description'])
end

def created_today?(x)
  Date.parse(x['created_at']).to_s == Date.today.to_s
end

def things_learned?(g)
  html_text = g['html_url']
  doc = Nokogiri::HTML(open(html_text))
  doc.xpath('//h1').count
end

@today_gist = @gists.select do |x|
  match_hacker_school_entries?(x) && created_today?(x)
end.first

@count = things_learned?(@today_gist)

Twitter.configure do |config|
  config.consumer_key = "uDqLBLtQNcbgB8GfcrmY6Q"
  config.consumer_secret = "9BlcQXCVNq9H6kLcaTySySRMIYjSwcGuotzbQtEM"
  config.oauth_token = "554640581-klhejXkw6yJulOWfdfkeAHJI8HEERNtcbV54KnUT"
  config.oauth_token_secret = "yR6cmygypsf4ycLWKhH5r6mpwHfw6i0QxSjMEFS7q8"
end

def catchphrase(count)
  random = rand(4)

  if count >= 5
    happy[random]
  elsif count.between?(3,5)
    meh[random]
  else
    sad[random]
  end
end

def happy
  [
    "Yippey Kai Yay Mother -!",
    "Eat my shorts"
  ]
end

def meh
  [
    "I'm Getting Too Old For This S***!",
    "Missed it by that much",
    "D'oh"
  ]
end

def sad
  [
    "Frankly Dear, I Don't Give A Damn",
    "Houston, We Have  A Problem.",
    "I've made a huge mistake",
    "God'll get you for that"
  ]
end

Twitter.update("learned #{@count} things #{@today_gist['html_url']} today. #{catchphrase(@count)}")
