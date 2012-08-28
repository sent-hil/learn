# call file with rescue pry_rescue.rb

def hi
  hola
end

hi

# or

Pry.rescue do
  hi
end
