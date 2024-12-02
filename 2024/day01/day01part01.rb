lhs = []
rhs = []
File.open('day01.txt', 'r') do |f|
  f.each_line do |line|
    split_line = line.split
    lhs << split_line[0].to_i
    rhs << split_line[1].to_i
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
