# BEGIN and END set up blocks that are called before anything else gets
# executed, or after everything else, just before the interpreter quits.
#
END { puts 'END block' }

puts 'foobar'

BEGIN { puts 'BEGIN block' }

# ruby this_file.rb outputs:
#
# BEGIN block
# foobar
# END block
