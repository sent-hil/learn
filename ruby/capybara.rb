require 'capybara/rspec'
require 'pry'

#Capybara.register_driver :chrome do |app|
  #Capybara::Selenium::Driver.new(app, :browser => :chrome)
#end

#Capybara.javascript_driver = :chrome

describe "Login ", js: true do
  include Capybara::DSL

  it do
    binding.pry
    response = visit('https://www.facebook.com/')
    fill_in :email, :with => ENV['USER']
    fill_in :pass, :with => ENV['PASS']
    click_button 'Log In'

    response = visit('https://www.facebook.com/dialog/oauth?response_type=code&client_id=369754433115833&redirect_uri=http%3A%2F%2Flocalhost%3A9393&scope=publish_stream')
 
    current_url.should == ''
  end
end
