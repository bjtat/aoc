require_relative '../utils/paths'

lhs = rhs = []
File.open(AocUtils.input_path(2024, 1), 'r') do |f|
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

lhs.zip(rhs).each_with_index do |left, right|
  sum += (left - right).abs
end

puts sum
