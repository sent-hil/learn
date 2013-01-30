# Pass custom flags to rspec via env variables.
#
# Run with `HI=1 rspec rspec_flags.rb`

RSpec.configure do |c|
  if ENV['HI']
    c.before(:all) do
      puts 'hi'
    end
  end
end

describe '' do
  it do
  end
end
