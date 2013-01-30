require 'parslet'
require 'pp'

class MiniP < Parslet::Parser
  rule(:lparen) { str('(') >> space? }
  rule(:rparen) { str(')') >> space? }
  rule(:comma)  { str(',') >> space? }

  rule(:equal)  { str('=') >> space? }

  rule(:space)   { match('\s').repeat(1) }
  rule(:space?)  { space.maybe }

  rule(:integer) { match('[0-9]').repeat(1).as(:int) >> space? }
  rule(:identifier) { match['a-z'].repeat(1) >> space? }
  rule(:operator) { match('[+]') >> space? }

  rule(:assignment) { identifier.as(:assignment) >> equal.as(:equal) >> expression }

  rule(:sum) { integer.as(:left) >> operator.as(:op) >> expression.as(:right) }
  rule(:arglist) { expression >> (comma >> expression).repeat}
  rule(:funcall) { identifier.as(:funcall) >> lparen >> arglist.as(:arglist) >> rparen }

  rule(:expression) { assignment | funcall | sum | integer }

  root(:expression)
end

pp MiniP.new.parse("a = 1")

class IntLit < Struct.new(:int)
  def eval
    int.to_i
  end
end

class Addition < Struct.new(:left, :right)
  def eval
    left.eval + right.eval
  end
end

class FunCall < Struct.new(:name, :args)
  def eval
    p args.map {|s| s.eval}
  end
end

class Assignment < Struct.new(:identifier, :right)
  def eval
    $identifier = right
  end
end

class MiniT < Parslet::Transform
  rule(:int => simple(:int)) do
    IntLit.new(int)
  end

  rule(:left => simple(:left), :right => simple(:right),
       :op => '+') do
         Addition.new(left, right)
  end

  rule(:funcall => 'puts', :arglist => subtree(:arglist)) do
    FunCall.new('puts', arglist)
  end

  rule(:variable => simple(:variable),
       :expression => simple(:expression)) do
    Assignment.new(variable, expression)
  end
end

#parser = MiniP.new
#transf = MiniT.new

#ast = transf.apply(
  #parser.parse(
    #'a = 1'))

#ast.eval
