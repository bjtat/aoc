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
rhs_ht = rhs.each_with_object(Hash.new(0)) { |num, counts| counts[num] += 1 }

similarity = 0
lhs.each do |left|
  similarity += left * rhs_ht[left]
end

puts similarity
