require 'pry'

# Kernel Extensions to support Lisp
class Object
  def lisp_string
    to_s
  end
end

class NilClass
  def lisp_string
    "nil" 
  end
end

class Array
  # Convert an Array into an S-expression (i.e. linked list).
  # Subarrays are converted as well.
  def sexp
    result = nil
    reverse.each do |item|
      item = item.sexp if item.respond_to?(:sexp)
      result = cons(item, result)
    end
    result
  end
end

# The Basic Lisp Cons cell data structures.  Cons cells consist of a
# head and a tail.
class Cons
  attr_reader :head, :tail

  def initialize(head, tail)
    @head, @tail = head, tail
  end

  def ==(other)
    return false unless other.class == Cons
    return true if self.object_id == other.object_id
    return car(self) == car(other) && cdr(self) == cdr(other)
  end

  # Convert the lisp expression to a string.
  def lisp_string
    e = self
    result = "(" 
    while e
      if e.class != Cons
        result << ". " << e.lisp_string
        e = nil
      else
        result << car(e).lisp_string
        e = cdr(e)
        result << " " if e
      end
    end
    result << ")" 
    result
  end
end

# Lisp Primitive Functions.

# It is an atom if it is not a cons cell.
def atom?(a)
  a.class != Cons
end

# Get the head of a list.
def car(e)
  e.head
end

# Get the tail of a list.
def cdr(e)
  e.tail
end

# Construct a new list from a head and a tail.
def cons(h,t)
  Cons.new(h,t)
end

# Here is the guts of the Lisp interpreter.  Apply and eval_one work
# together to interpret the S-expression.  These definitions are taken
# directly from page 13 of the Lisp 1.5 Programmer's Manual.

def apply(fn, x, a)
  if atom?(fn)
    case fn
    when :car then caar(x)
    when :cdr then cdar(x)
    when :cons then cons(car(x), cadr(x))
    when :atom then atom?(car(x))
    when :eq then car(x) == cadr(x)
    else
      apply(eval_one(fn,a), x, a)
    end
  elsif car(fn) == :lambda
    eval_one(caddr(fn), pairlis(cadr(fn), x, a))
  elsif car(fn) == :label
    apply(caddr(fn), x, cons(cons(cadr(fn), caddr(fn)), a))
  end
end

def eval_one(e,a)
  binding.pry

  if atom?(e)
    cdr(assoc(e,a))
  elsif atom?(car(e))
    if car(e) == :quote
      cadr(e)
    elsif car(e) == :cond
      evcon(cdr(e),a)
    else
      apply(car(e), evlis(cdr(e), a), a)
    end
  else
    apply(car(e), evlis(cdr(e), a), a)
  end
end

# And now some utility functions used by apply and eval_one.  These are
# also given in the Lisp 1.5 Programmer's Manual.

def evcon(c,a)
  if eval_one(caar(c), a)
    eval_one(cadar(c), a)
  else
    evcon(cdr(c), a)
  end
end

def evlis(m, a)
  if m.nil?
    nil
  else
    cons(eval_one(car(m),a), evlis(cdr(m), a))
  end
end

def assoc(a, e)
  if e.nil?
    fail "#{a.inspect} not bound" 
  elsif a == caar(e)
    car(e)
  else
    assoc(a, cdr(e))
  end
end

def pairlis(vars, vals, a)
  while vars && vals
    a = cons(cons(car(vars), car(vals)), a)
    vars = cdr(vars)
    vals = cdr(vals)
  end
  a
end

# Handy lisp utility functions built on car and cdr.

def caar(e)
  car(car(e))
end

def cadr(e)
  car(cdr(e))
end

def caddr(e)
  car(cdr(cdr(e)))
end

def cdar(e)
  cdr(car(e))
end

def cadar(e)
  car(cdr(car(e)))
end

env = [
  cons(:rev_shift,
    [:lambda, [:list, :result],
      [:cond,
        [[:null, :list], :result],
        [:t, [:rev_shift, [:cdr, :list],
            [:cons, [:car, :list], :result]]]]].sexp),
  cons(:reverse,
    [:lambda, [:list], [:rev_shift, :list, nil]].sexp),
  cons(:null, [:lambda, [:e], [:eq, :e, nil]].sexp),
  cons(:t, true), 
  cons(nil, nil)
].sexp

# eval_oneuate an S-Expression and print the result

exp = [:reverse, [:quote, [:a, :b, :c, :d, :e]]].sexp

puts "eval_one: #{exp.lisp_string}" 
puts "  =>  #{eval_one(exp,env).lisp_string}"
