b = 1

class Test
  def test
    @a = 1
    b = 2

    [@a,b]
  end
end

set_trace_func proc { |event, file, line, id, binding, classname|
  instance_vars = binding.eval("instance_variables").inject({}) do |vars, name|
    value = binding.eval(name.to_s)
    vars[name] = value unless value.nil?
    vars
  end

  local_vars = binding.eval("local_variables").inject({}) do |vars, name|
    value = binding.eval(name.to_s)
    vars[name] = value unless value.nil?
    vars
  end

  printf "%8s %-2d %10s %8s %10s %10s\n", event, line, id, classname, local_vars, instance_vars
}
t = Test.new
t.test
