require 'rubygems'
require 'twitter'
require 'pry'

Twitter.configure do |config|
  config.consumer_key = "ceBCZQQai1QT1vwzbAba4w"
  config.consumer_secret = "gRWFq4EizdyqDnQj76Wdq7oLJONiosB0HER1qASOk"
  config.oauth_token = "554640581-5ATVXqll3QuKYeNMI5BH1GP8dEQBIGSELrHxT5hF"
  config.oauth_token_secret = "4LQYqQ5rDow6FUE2cP73SZQ1c6XZm2lI38B0UNh8"
end

binding.pry

def get_twitter_friends_with_cursor(cursor, list, client)
  # Base case
  if cursor == 0
    return list
  else
    hashie = client.friends(:cursor => cursor)
    hashie.users.each {|u| list << u } # Concat users to list
    get_twitter_friends_with_cursor(hashie.next_cursor,list,client) # Recursive step using the next cursor
  end
end

# Generate Oauth
user = User.find.first
oauth = Twitter::OAuth.new(TOKEN,SECRET)
oauth.authorize_from_access(user.twitter_token,user.twitter_secret)
client = Twitter::Base.new(oauth)

# recursively get all Twitter friends
# pass the default -1 cursor, an empty list and the Oauth client
all_friends = get_twitter_friends_with_cursor(-1,[],client)
