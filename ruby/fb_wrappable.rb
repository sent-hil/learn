module Facebook
  include ApiWrappable

  endpoint "https://graph.facebook.com"

  responds_to :xml

  get "me" do
    get "friends"
    get "feed"
    get "movies"
    get "books"
    get "likes"

    post "feed"
  end

  get ":id" do
    get "picture"
  end

  def before_request(request)
    request.params['access_token'] = 'AAA'
  end

  def after_request(response)
    JSON.parse(response)
  end
end

Facebook::Id.new(1, :access_token => '').get

fb_id = Facebook::Id.new(1)
fb_id.get
fb_id.pictures.get

me = Facebook::Me.new
me.get
me.friends.get
me.feed.get
me.movies.get
me.books.get
me.likes.get

me.feed(:message => 'Hello!').post

me.likes.limit(3).start('1').get
me.likes(limit:3, start:1).get
