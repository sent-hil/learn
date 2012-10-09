module Twitter
  include ApiWrappable

  endpoint "https://api.twitter.com/1.1"

  responds_to :json

  get "statuses" do
    get "user_timeline"
    get "home_timeline"

    get "retweets/:id"
    get "show/:id"

    post "destroy/:id"
  end

  def before_post_request(request)
    request.headers['Accept'] = 'application/json'
  end
end

me = Twitter::Statuses.new(access_token:'AAA')
me.user_timeline.get
me.home_timeline.get
me.retweets(1).get
me.show(1).get
me.destroy(1).post
