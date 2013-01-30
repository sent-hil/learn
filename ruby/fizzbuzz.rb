def fizzbuzz
  (1..100).to_a.each do |x|
    three_divisble = x % 3 == 0
    five_divisble  = x % 5 == 0
    three_and_five_divisible = three_divisble && five_divisble

    binding.pry

    if three_and_five_divisible
      puts "FizzBuzz"
    elsif three_divisble
      puts "Fizz"
    elsif five_divisble
      puts "Buzz"
    else
      puts x
    end
  end
end

fizzbuzz
