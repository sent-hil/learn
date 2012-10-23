set_trace_func proc { |event, file, line, id, bind, classname|
  locals = bind.eval('local_variables').count
  #puts locals

  printf "%8s %s:%-2d %10s %8s %12s\n", event, file, line, id, classname, locals
}

one = 'hello'
two = 'hola'
three = two

#class Test
  #def test
    #a = 1
  #end

  #def second
  #end
#end

#t = Test.new
#t.test
