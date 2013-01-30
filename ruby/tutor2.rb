$local_values_at_time = {}
$instance_values_at_time = {}

set_trace_func proc { |event, file, line, id, binding, classname|
  $local_values_at_time[line] = binding.eval('local_variables.map {|v| [v, eval(v.to_s).object_id, v.object_id] }')
  #$instance_values_at_time[line] = binding.eval('instance_variables.map {|v| {v => eval(v.to_s)} }')
}

#@one = 0
#x = nil
#a = 5
#stuff = 'asd'
#a += 8

foo = 1
b = [foo,2,3]
z = b[0]

set_trace_func nil
$local_values_at_time.each do |key, pair|
  puts "#{key}=>#{pair}"
end

puts

$instance_values_at_time.each do |key, pair|
  puts "#{key}=>#{pair}"
end

#trace_var :$hello, proc {|v| puts "$_ is now '#{v}'" }

#$hello = "hello"

#class One
  #def self.singleton_method_added(id)
    #puts "Added #{id} to One"
  #end

  #def a
  #end

  #def One.b
    #puts 'b'
  #end

  #self.class.send(:define_method, :c, proc { puts 'c' })
#end

#One.b
