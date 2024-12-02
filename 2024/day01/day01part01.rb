lhs = rhs = []
File.open('day01.txt', 'r') do |f|
  f.each_line do |line|
    left, right = line.split.map(&:to_i)
    lhs << left.to_i
    rhs << right.to_i
  end
  f.close
end

lhs.sort!
rhs.sort!

sum = 0
lhs.zip(rhs).each do |left, right|
  sum += (left - right).abs
end

puts sum
