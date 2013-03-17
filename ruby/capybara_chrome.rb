#Capybara.register_driver :chrome do |app|
  #Capybara::Selenium::Driver.new(app, :browser => :chrome)
#end
#Capybara.javascript_driver = :chrome

#RSpec.configure do |c|
  #c.around(:each, :js => true) do |example|
    #if Capybara.current_driver == :chrome
      #require 'headless'

      #headless = Headless.new
      #headless.start
    #end
    #example.run
    #headless.destroy if Capybara.current_driver == :chrome
  #end
#end
