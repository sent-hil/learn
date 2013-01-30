require 'ostruct'

Config = Struct.new(:name)
c = Config.new
c.name = 'indiana'
puts c.name

cf = OpenStruct.new(:name => 'jones')
puts cf.name
