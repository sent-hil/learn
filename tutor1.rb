#class Test
  #def test
    #@a = 1
    #b = 2

    #[@a,b]
  #end
#end

set_trace_func proc { |event, file, line, id, binding, classname|
  #instance_vars = binding.eval("instance_variables").inject({}) do |vars, name|
    #value = binding.eval(name.to_s)
    #vars[name] = value unless value.nil?
    #vars
  #end

  local_vars = binding.eval("local_variables").inject({}) do |vars, name|
    value = binding.eval(name.to_s)
    vars[name] = value
    vars
  end

  puts local_vars

  printf "%8s %-2d %10s %8s %10s \n", event, line, id, classname, local_vars
}

@a = 1
set_trace_func nil
puts binding.eval("instance_variables")
@b = 1
puts binding.eval("@b")
c = nil
