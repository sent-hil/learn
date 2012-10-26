require 'pry'

$local = {}

set_trace_func proc { |event, file, line, id, bind, classname|
  bind.eval('local_variables').each do |l|
    $local[l] - bind.eval(l)
  end

  printf "%8s %s:%-2d %10s %8s %12s\n", event, file, line, id, classname, locals
}

one = 'hello'
two = 'hola'
#three = two

#class Test
  #def test
    #a = 1
  #end

  #def second
  #end
#end

#t = Test.new
#t.test

set_trace_func nil
